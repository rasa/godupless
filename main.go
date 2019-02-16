// Program godupless creates a report of duplicate files across multiple volumes
package main

import (
	"crypto/md5"
	"crypto/sha1"
	"crypto/sha256"
	"crypto/sha512"
	"encoding/hex"
	"flag"
	"fmt"
	"hash"
	"io"
	"os"
	"path/filepath"
	"regexp"
	"runtime"
	"sort"
	// "strconv"
	"strings"
	"time"

	"github.com/cespare/xxhash"
	"github.com/minio/highwayhash"
	"github.com/rasa/godupless/file"
	"github.com/rasa/godupless/util"
	"github.com/rasa/godupless/version"
	"golang.org/x/text/language"
	"golang.org/x/text/message"
)

const (
	MinChunk = 2 << 12 // 4096
	MaxChunk = 2 << 24// 16,777,216
)

// Dir @todo
type Dir struct {
	// Count @todo
	Count uint64
	// Size @todo
	Size uint64
}

// DirCount @todo
type DirCount struct {
	// Dir contains the full of the directory containing one or more files that are duplicated
	Dir string
	// Count contains the number of files in directory that are duplicated in other directories
	Count uint64
}

// ErrorRec @todo
type ErrorRec struct {
	// Path contains the full path of the file that generated an error
	Path string
	// Error contains the error message related to the file
	Error string
}

// IgnoredRec @todo
type IgnoredRec struct {
	// Path contains the full path of the ignore file
	Path string
	// Type contains the type of ignored file (symlink, named pipe, etc)
	Type string
}

// Config @todo
type Config struct {
	//cache     string
	chunk     uint
	//separator string
	exclude   string
	//extra         bool
	freq     uint
	hash     string
	help     bool
	iexclude string
	mask     string
	// @todo add and imask option?
	minDirLength uint
	minFiles     uint
	minSize      uint64
	recursive    bool
	// utc          bool
	verbose uint

	dirReport     bool
	errorReport   bool
	ignoredReport bool
	sizeReport    bool
}

var config = Config{
	//cache: "godupless.cache",
	chunk: 2 << 19, // 2<<19=2^20=1,048,576
	//separator: ",",
	//exclude: "",
	//extra: false,
	freq: 100,
	hash: "highway",
	//help: false,
	//iexclude: "",
	//mask: "",
	minDirLength: 2,
	minFiles: 2,
	minSize: 2 << 20, // 2<<20=2^21=2,097,152
	//recursive: false,
	// utc: false,
	//verbose: 0,

	//dirReport: false,
	//errorReport: false,
	//ignoredReport: false,
	sizeReport: true,
}

// Stats @todo
type Stats struct {
	// device stats:
	hits        uint
	skipped     uint
	directories uint
	matched     uint
	errors      uint
	ignored     uint
}

// Dupless @todo
type Dupless struct {
	config   Config
	stats    Stats
	excludes []string
	masks    []string
	p        *message.Printer
	path     string
	dev      string
	lastDev  string
	volume   string

	cacheFH *os.File
	// comma   rune

	// dirs[dir] = *Dir
	// errorDirs[dir] = *ErrorRec[]
	// ignoredDirs[dir] = *IgnoredRec[]
	// files[path] = *file.File
	// uniques[uniqueID] = paths[]
	// sizes[size][uniqueIDs] = paths[]
	// hashes[size][hash] = *file.File[]
	dirs        map[string]*Dir
	errorDirs   map[string][]*ErrorRec
	ignoredDirs map[string][]*IgnoredRec
	files       map[string]*file.File
	uniques     map[string][]string
	sizes       map[uint64]map[string][]*file.File
	hashes      map[uint64]map[string][]*file.File
}

// Uint64Slice @todo
type Uint64Slice []uint64

func (p Uint64Slice) Len() int           { return len(p) }
func (p Uint64Slice) Less(i, j int) bool { return p[i] < p[j] }
func (p Uint64Slice) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }

// Sort is a convenience method.
func (p Uint64Slice) Sort() { sort.Sort(p) }

// utility functions

func (d *Dupless) init() {
	d.config = config
	minDirLength := d.config.minDirLength
	if runtime.GOOS == "windows" {
		// c:/
		minDirLength = 3
	}

	// flag.StringVar(&d.config.cache, "cache", DefaultCache, "Cache filename")
	flag.UintVar(&d.config.chunk, "chunk", d.config.chunk, "Hash chunk")
	flag.BoolVar(&d.config.dirReport, "dir_report", d.config.dirReport, "Report by directory")
	flag.BoolVar(&d.config.errorReport, "error_report", d.config.errorReport, "Report of errors")
	flag.StringVar(&d.config.exclude, "exclude", d.config.exclude, "Regex(s) of Directories/files to exclude, separated by |")
	flag.StringVar(&d.config.iexclude, "iexclude", d.config.iexclude, "Regex(s) of Directories/files to exclude, separated by |")
	//flag.BoolVar(&d.config.extra, "extra", d.config.extra, "Cache extra attributes")
	flag.UintVar(&d.config.freq, "frequency", d.config.freq, "Reporting frequency")
	flag.StringVar(&d.config.hash, "hash", d.config.hash, "Hash type")
	flag.BoolVar(&d.config.help, "help", d.config.help, "Display help")
	flag.BoolVar(&d.config.ignoredReport, "ignored_report", d.config.ignoredReport, "Report of ignored files")
	flag.StringVar(&d.config.mask, "mask", d.config.mask, "File mask(s), seperated by |")
	flag.UintVar(&d.config.minDirLength, "min_dir_len", minDirLength, "Minimum directory length")
	flag.UintVar(&d.config.minFiles, "min_files", d.config.minFiles, "Minimum files")
	flag.Uint64Var(&d.config.minSize, "min_size", d.config.minSize, "Minimum file size")
	flag.BoolVar(&d.config.recursive, "recursive", d.config.recursive, "Report directories recursively")
	flag.BoolVar(&d.config.sizeReport, "size_report", d.config.sizeReport, "Report by size")
	//flag.StringVar(&d.config.seperator, "separator", d.config.seperator, "Field separator")
	// flag.BoolVar(&d.config.utc, "utc", d.config.utc, "Report times in UTC")
	flag.UintVar(&d.config.verbose, "verbose", d.config.verbose, "Be more verbose")

	flag.Parse()
	
	if d.config.chunk < MinChunk || d.config.chunk > MaxChunk {
		fmt.Printf("Chunk must be between %d and %d: %d", MinChunk, MaxChunk, d.config.chunk)
		os.Exit(1)
	}

	d.config.hash = strings.ToLower(d.config.hash)

	if runtime.GOOS != "windows" {
		d.excludes = []string{
			"^/dev$",
			"^/proc$",
			"^/run$",
			"^/sys$",
		}
	} else {
		// @todo ignore all hidden/system directories?
		d.excludes = []string{
			`(?i)^[A-Z]:/$Recycle\.bin`,
			`(?i)^[A-Z]:/System Volume Information`,
		}
	}

	if d.config.exclude != "" {
		a := strings.Split(d.config.exclude, "|")
		for _, s := range a {
			s = util.NormalizePath(s)
			d.excludes = append(d.excludes, s)
		}
	}

	if d.config.iexclude != "" {
		a := strings.Split(d.config.iexclude, "|")
		for _, s := range a {
			s = util.NormalizePath(s)
			d.excludes = append(d.excludes, "(?i)"+s)
		}
	}

	//util.Dump("d.excludes=", d.excludes)

	if d.config.mask != "" {
		a := strings.Split(d.config.mask, "|")
		for _, s := range a {
			d.masks = append(d.masks, s)
		}
	}

	// value, _ /*multibyte*/, _ /*tail*/, err := strconv.UnquoteChar(d.config.separator, 0)
	/*
		if err != nil {
			panic(err)
		}
		d.comma = value
	*/

	d.dirs = make(map[string]*Dir)
	d.errorDirs = make(map[string][]*ErrorRec)
	d.ignoredDirs = make(map[string][]*IgnoredRec)

	d.files = make(map[string]*file.File)
	d.uniques = make(map[string][]string)
	d.sizes = make(map[uint64]map[string][]*file.File)

	// @todo determine language from OS
	d.p = message.NewPrinter(language.English)
}

func (d *Dupless) addPath(path string, size uint64) {
	for {
		dir := util.Dirname(path)
		if uint(len(dir)) < d.config.minDirLength {
			return
		}
		if dir == path {
			return
		}
		_, ok := d.dirs[dir]
		if !ok {
			d.dirs[dir] = &Dir{Count: 1, Size: size}
		} else {
			d.dirs[dir].Count++
			d.dirs[dir].Size += size
		}
		if !d.config.recursive {
			return
		}
		path = dir
	}
}

func (d *Dupless) progress(final bool) {
	if !final && (d.config.freq == 0 || d.stats.hits%d.config.freq != 0) {
		return
	}

	dev := d.dev
	if d.volume > "" {
		dev += " (" + d.volume + ")"
	}
	d.p.Printf("\r%11d %11d %11d %11d %11d %s", d.stats.skipped, d.stats.matched, d.stats.directories, d.stats.ignored, d.stats.errors, dev)

	if final {
		fmt.Println("")
	}
}

func (d *Dupless) addError(path string, s string) {
	dir := util.Dirname(path)
	errorRec := ErrorRec{Path: path, Error: s}
	d.errorDirs[dir] = append(d.errorDirs[dir], &errorRec)
	d.stats.errors++
	if d.config.verbose > 0 {
		fmt.Fprintf(os.Stderr, "\n%s\n", s)
	}
}

func (d *Dupless) addIgnore(path string, typ string) {
	dir := util.Dirname(path)
	IgnoredRec := IgnoredRec{Path: path, Type: typ}
	d.ignoredDirs[dir] = append(d.ignoredDirs[dir], &IgnoredRec)
	d.stats.ignored++
	if d.config.verbose > 0 {
		fmt.Fprintf(os.Stderr, "\nSkipping '%s': %s\n", path, typ)
	}
}

func (d *Dupless) reportByDir() {
	fmt.Printf("\nDuplication Report By Size/Directory\n\n")

	for size, hashmap := range d.hashes {
		for _, files := range hashmap {
			for _, f := range files {
				d.addPath(f.Path(), size)
			}
		}
	}

	sizemap := make(map[uint64][]*DirCount)

	for dir, d := range d.dirs {
		size := d.Size
		/*_, ok := sizemap[size]
		if !ok {
			sizemap[size] = make([]*DirCount, 0)
		}*/
		sizemap[size] = append(sizemap[size], &DirCount{Dir: dir, Count: d.Count})
	}

	i := 0
	sizes := make([]uint64, len(sizemap))
	for size := range sizemap {
		sizes[i] = size
		i++
	}
	sort.Sort(sort.Reverse(Uint64Slice(sizes)))

	var totalSize uint64
	var files uint64

	for _, size := range sizes {
		emitSize := true
		totalSize += size
		for i, dircount := range sizemap[size] {
			if emitSize {
				fmt.Printf("size: %d\n", size)
				emitSize = false
			}
			d.p.Printf("  %d: %v (%d files)\n", i+1, dircount.Dir, dircount.Count)
			files += dircount.Count
		}
	}
	d.p.Printf("\n%d files totaling %d bytes (%d bytes per file average)\n", files, totalSize, totalSize/files)
}

func (d *Dupless) reportBySize() {
	fmt.Printf("\nDuplication Report By Size/Paths\n\n")

	i := 0
	sizes := make([]uint64, len(d.hashes))
	for size := range d.hashes {
		sizes[i] = size
		i++
	}
	sort.Sort(sort.Reverse(Uint64Slice(sizes)))

	var totalSize uint64
	var totalFiles uint64

	for _, size := range sizes {
		for hash, files := range d.hashes[size] {
			totalFiles += uint64(len(files))
			d.p.Printf("Size: %d (%s)\n", size, hash)
			for i, f := range files {
				totalSize += size
				var s string
				if f.Links() > 1 {
					s = fmt.Sprintf(" (%d hard links)", f.Links())
				}
				fmt.Printf("  %d: %s%s\n", i+1, f.Path(), s)
			}
		}
	}

	var avg uint64
	if totalFiles > 0 {
		avg = totalSize / totalFiles
	}

	d.p.Printf("\n%d files totaling %d bytes (%d bytes per file average)\n", totalFiles, totalSize, avg)
}

func (d *Dupless) reportIgnored() {
	fmt.Printf("\nIgnored Files/Directories Report\n\n")

	i := 0
	dirs := make([]string, len(d.ignoredDirs))
	for dir := range d.ignoredDirs {
		dirs[i] = dir
		i++
	}
	sort.Strings(dirs)

	for _, dir := range dirs {
		for _, other := range d.ignoredDirs[dir] {
			fmt.Printf("%-10s: %s\n", other.Type, other.Path)
		}
	}
}

func (d *Dupless) reportErrors() {
	fmt.Printf("\nError Files/Directories Report\n\n")

	i := 0
	dirs := make([]string, len(d.errorDirs))
	for dir := range d.errorDirs {
		dirs[i] = dir
		i++
	}
	sort.Strings(dirs)

	for _, dir := range dirs {
		for _, other := range d.errorDirs[dir] {
			fmt.Println(other.Error)
		}
	}
}

func (d *Dupless) countFiles(hashes map[uint64]map[string][]*file.File) (scanning int, total int) {
	scanning = 0
	total = 0
	for _, hashmap := range hashes {
		for _, files := range hashmap {
			for _, f := range files {
				if f.Err() == nil {
					total++
				}
				if !f.EOF() {
					scanning++
				}
			}
		}
	}
	return scanning, total
}

func (d *Dupless) getHash(hashes map[uint64]map[string][]*file.File) error {
	eof := true
	for _, hashmap := range hashes {
		for _, files := range hashmap {
			for _, f := range files {
				if f.EOF() {
					continue
				}
				if f.Err() != nil {
					continue
				}
				if !f.Opened() {
					//fmt.Printf("Opening %s\n", f.Path())
					err := f.Open()
					if err != nil {
						fmt.Printf("doHash: %s\n", err)
						continue
					}
				}
				//fmt.Printf("Reading %d bytes from %s\n", d.chunk, f.Path())
				err := f.Read(uint64(d.config.chunk))
				if err == nil {
					eof = false
					continue
				}
				if err != io.EOF {
					// don't set eof to false on errors
					fmt.Printf("dohash: %s\n", err)
				}
				err = f.Close()
				if err != nil {
					fmt.Printf("dohash: %s\n", err)
				}
			}
		}
	}
	if eof {
		return io.EOF
	}
	return nil
}

func (d *Dupless) regenHashTable(hashes map[uint64]map[string][]*file.File) (newHashes map[uint64]map[string][]*file.File) {
	newHashes = make(map[uint64]map[string][]*file.File)

	for size, hashmap := range hashes {
		_, ok := newHashes[size]
		if !ok {
			newHashes[size] = make(map[string][]*file.File)
		}
		for _, files := range hashmap {
			for _, f := range files {
				if f.Err() != nil {
					//fmt.Printf("regenHashTable(): error: %s\n", f.Err())
					continue
				}
				_, ok = newHashes[size][f.Hash()]
				if !ok {
					newHashes[size][f.Hash()] = make([]*file.File, 0)
				}
				newHashes[size][f.Hash()] = append(newHashes[size][f.Hash()], f)
			}
		}
	}

	for size, hashmap := range newHashes {
		for hash, files := range hashmap {
			if uint(len(files)) < d.config.minFiles {
				for _, f := range files {
					f.Close()
				}
				//log.Printf("Deleting size %d hash %s\n", size, hash)
				delete(newHashes[size], hash)
			}
		}
	}

	for size, hashmap := range newHashes {
		if len(hashmap) < 1 {
			//log.Printf("Deleting size %d\n", size)
			delete(newHashes, size)
		}
	}

	return newHashes
}

func (d *Dupless) getHasher() hash.Hash {
	// @todo move to constant
	skey := "0000000000000000000000000000000000000000000000000000000000000000"
	key, _ := hex.DecodeString(skey)

	switch d.config.hash {
	case "highway64":
		h, _ := highwayhash.New64(key)
		return h
	case "highway128":
		h, _ := highwayhash.New128(key)
		return h
	case "highway256", "highway":
		h, _ := highwayhash.New(key)
		return h
	case "md5":
		h := md5.New()
		return h
	case "sha1":
		h := sha1.New()
		return h
	case "sha256":
		h := sha256.New()
		return h
	case "sha512":
		h := sha512.New()
		return h
	case "xxhash":
		h := xxhash.New()
		return h
	default:
		fmt.Fprintf(os.Stderr, "\nUnknown hash format: '%s'\n", d.config.hash)
		os.Exit(1)
	}

	return nil
}

func (d *Dupless) getHashes() bool {
	for path, f := range d.files {
		_, ok := d.uniques[f.UniqueID()]
		if !ok {
			d.uniques[f.UniqueID()] = make([]string, 0)
		}
		d.uniques[f.UniqueID()] = append(d.uniques[f.UniqueID()], path)
		_, ok = d.sizes[f.Size()]
		if !ok {
			d.sizes[f.Size()] = make(map[string][]*file.File)
		}
		_, ok = d.sizes[f.Size()][f.UniqueID()]
		if !ok {
			d.sizes[f.Size()][f.UniqueID()] = make([]*file.File, 0)
		}
		d.sizes[f.Size()][f.UniqueID()] = append(d.sizes[f.Size()][f.UniqueID()], f)
	}

	if len(d.sizes) < 1 {
		return false
	}

	//util.Dump("d.files=", d.files)
	//util.Dump("d.uniques=", d.uniques)

	for size, uniques := range d.sizes {
		if uint(len(uniques)) < d.config.minFiles {
			delete(d.sizes, size)
		}
	}
	//util.Dump("d.sizes=", d.sizes)

	if len(d.sizes) < 1 {
		return false
	}

	var sizes = make([]uint64, len(d.sizes))

	i := 0
	for size := range d.sizes {
		sizes[i] = size
		i++
	}

	sort.Sort(sort.Reverse(Uint64Slice(sizes)))

	h := d.getHasher()

	defaultHash := fmt.Sprintf("%x", h.Sum(nil))

	var hashes = make(map[uint64]map[string][]*file.File)
	for size, uniqueMap := range d.sizes {
		_, ok := hashes[size]
		if !ok {
			hashes[size] = make(map[string][]*file.File)
		}
		_, ok = hashes[size][defaultHash]
		if !ok {
			hashes[size][defaultHash] = make([]*file.File, 0)
		}
		for _, files := range uniqueMap {
			// we only need to hash the first file, as the others are the same file
			hashes[size][defaultHash] = append(hashes[size][defaultHash], files[0])
		}
	}

	//util.Dump("hashes=", hashes)

	loops := 1 + (sizes[0] / uint64(d.config.chunk))

	loop := 0
	read := uint64(0)
	for {
		loop++
		scanning, total := d.countFiles(hashes)
		d.p.Printf("Loop %d of %d: %d of %d bytes read: scanning %d of %d files (%d unique sizes)\n", loop, loops, read, sizes[0], scanning, total, len(hashes))
		read += uint64(d.config.chunk)
		err := d.getHash(hashes)
		if err == io.EOF {
			fmt.Println("All files have been hashed")
			break
		}
		if err != nil {
			fmt.Println(err)
		}
		newHashes := d.regenHashTable(hashes)
		hashes = newHashes
		if len(hashes) == 0 {
			break
		}
	}

	//elapsed := time.Since(start)
	//_, total := d.countFiles(hashes)
	//fmt.Printf("\nHashed %d files in %s\n", total, elapsed)
	//util.Pause()
	//util.Dump("hashes=", hashes)
	d.hashes = hashes

	if len(d.hashes) < 1 {
		return false
	}

	return true
}

func (d *Dupless) visit(path string, fi os.FileInfo, err error) error {
	path = util.NormalizePath(path)
	if d.config.verbose > 0 {
		fmt.Printf("Opening %s\n", path)
	}
	if err != nil {
		if d.config.verbose > 0 {
			fmt.Fprintf(os.Stderr, "\nError on '%s': %s\n", path, err)
		}
	}

	err = nil
	for {
		d.path = path
		d.stats.hits++

		for _, exclude := range d.excludes {
			ok, e := regexp.MatchString(exclude, path)
			if e != nil {
				s := fmt.Sprintf("Failed to match '%s' via '%s': %s", path, exclude, e)
				d.addError(path, s)
			}
			if ok {
				d.addIgnore(path, "excluded")
				return filepath.SkipDir
			}
		}

		if len(d.masks) > 0 {
			//_, file := filepath.Split(path)
			matched := false
			for _, mask := range d.masks {
				ok, e := filepath.Match(mask, path)
				if e != nil {
					panic(e)
				}
				if ok {
					matched = true
					break
				}
			}
			if !matched {
				d.stats.skipped++
				break
			}
		}

		if fi != nil {
			if fi.IsDir() {
				d.stats.directories++
				break
			}
			if fi.Mode()&os.ModeSymlink != 0 {
				d.stats.ignored++
				break
			}
			if !fi.Mode().IsRegular() {
				d.addIgnore(path, fi.Mode().String())
				break
			}

			if uint64(fi.Size()) <= d.config.minSize {
				d.stats.skipped++
				break
			}
		}

		f, e := file.NewFile(path, fi, d.getHasher())
		if e != nil {
			s := fmt.Sprintf("Cannot stat '%s': %s", path, e)
			d.addError(path, s)
			break
		}

		if f.IsDir() {
			d.stats.directories++
			break
		}
		if f.IsSymlink() {
			d.stats.ignored++
			break
		}
		if !f.IsRegular() {
			d.addIgnore(path, f.Type())
			break
		}

		if f.Size() <= d.config.minSize {
			d.stats.skipped++
			break
		}

		d.dev, _ = f.VolumeName()
		if d.lastDev != d.dev {
			// @todo add option to include/exclude cross-device files
			if d.lastDev != "" {
				// @todo log skipped files
				fmt.Printf("\nSkipping %s as it is on device %s\n", path, d.dev)
				return filepath.SkipDir
			}
			d.lastDev = d.dev
		}

		d.stats.matched++
		d.files[path] = f
		break
	}

	d.progress(false)
	return err
}

func (d *Dupless) resetCounters() {
	d.stats.skipped = 0
	d.stats.directories = 0
	d.stats.matched = 0
	d.stats.errors = 0
	d.stats.ignored = 0
	d.lastDev = ""
}

func (d *Dupless) progressHeader() {
	fmt.Println("")
	d.p.Printf("Chunk size:    %d\n", d.config.chunk)
	d.p.Printf("Minimum files: %d\n", d.config.minFiles)
	d.p.Printf("Minimum size:  %d\n", d.config.minSize)
	d.p.Printf("Masks:         %v\n", d.config.mask)
	d.p.Printf("Recursive:     %v\n", d.config.recursive)
	d.p.Printf("Verbosity:     %d\n", d.config.verbose)
	fmt.Println("")

	fmt.Printf("    skipped     matched directories     ignored      errors device\n")
	fmt.Printf("----------- ----------- ----------- ----------- ----------- ------\n")
}

func (d *Dupless) findFiles() bool {
	var args []string
	for _, arg := range flag.Args() {
		if arg == "*" {
			volumes, err := file.GetVolumes()
			if err != nil {
				fmt.Fprintln(os.Stderr, "\nFailed to get volumes:", err)
			}
			for _, volume := range volumes {
				args = append(args, volume)
			}
			continue
		}
		if runtime.GOOS == "windows" {
			if len(arg) == 2 && arg[1] == ':' {
				arg += string(os.PathSeparator)
			}
		}
		args = append(args, arg)
	}

	for i, arg := range args {
		if i > 0 {
			fmt.Println("")
		}
		d.resetCounters()
		if file.IsVolume(arg) {
			d.volume = arg
		} else {
			d.volume = ""
		}
		err := filepath.Walk(arg, d.visit)
		if err != nil {
			fmt.Fprintln(os.Stderr, "\nWalk returned:", err)
		}
	}

	d.progress(true)

	if len(d.files) < 1 {
		return false
	}
	return true
}

func (d *Dupless) doReports() {
	if d.config.errorReport {
		d.reportErrors()
	}
	if d.config.ignoredReport {
		d.reportIgnored()
	}
	if d.config.dirReport {
		d.reportByDir()
	}
	if d.config.sizeReport {
		d.reportBySize()
	}
}

func (d *Dupless) footer(elapsed time.Duration, elapsed2 time.Duration) {
	fmt.Printf("\nFound %d matching files in %s\n", len(d.files), elapsed)
	_, total := d.countFiles(d.hashes)
	fmt.Printf("Hashed %d files in %s\n", total, elapsed2)
	fmt.Printf("Total elapsed time: %s\n", elapsed+elapsed2)
}

func usage() {
	fmt.Fprintf(os.Stderr, "\nUsage: %s [options] path [path2] ...\nOptions:\n\n", os.Args[0])
	flag.PrintDefaults()
}

func main() {
	basename := filepath.Base(os.Args[0])
	progname := strings.TrimSuffix(basename, filepath.Ext(basename))

	var d Dupless

	d.init()

	fmt.Printf("%s: Version %s (%s)\n", progname, version.VERSION, version.GITCOMMIT)
	fmt.Printf("Built with %s for %s/%s (%d CPUs/%d GOMAXPROCS)\n",
		runtime.Version(),
		runtime.GOOS,
		runtime.GOARCH,
		runtime.NumCPU(),
		runtime.GOMAXPROCS(-1))

	if len(flag.Args()) == 0 || d.config.help {
		usage()
		return
	}

	d.progressHeader()

	start := time.Now()

	if !d.findFiles() {
		fmt.Println("No matching files found")
		return
	}

	elapsed := time.Since(start)

	start2 := time.Now()

	ok := d.getHashes()

	elapsed2 := time.Since(start2)

	if ok {
		d.doReports()
	} else {
		fmt.Printf("No duplicate files found\n")
	}

	d.footer(elapsed, elapsed2)
}
