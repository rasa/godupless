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
	"strconv"
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
	// DefaultCache @todo
	DefaultCache = "godupless.cache"
	// DefaultChunk @todo
	DefaultChunk = 2 << 19 // 2<<19=2^20=1,048,576
	// DefaultDirReport @todo
	DefaultDirReport = false
	// DefaultExclude @todo
	DefaultExclude = ""
	// DefaultExtra @todo
	DefaultExtra = false
	// DefaultFrequency @todo
	DefaultFrequency = 100
	// DefaultHash @todo
	DefaultHash = "highway"
	// DefaultIexclude @todo
	DefaultIexclude = ""
	// DefaultMask @todo
	DefaultMask = ""
	// DefaultMinFiles @todo
	DefaultMinFiles = 2
	// DefaultMinDirLength @todo
	DefaultMinDirLength = uint(2)
	// DefaultMinSize @todo
	DefaultMinSize = 2 << 20 // 2<<20=2^21=2,097,152
	// DefaultRecursive @todo
	DefaultRecursive = false
	// DefaultSeparator @todo
	DefaultSeparator = ","
	// DefaultSizeReport @todo
	DefaultSizeReport = true
	// DefaultUTC @todo
	DefaultUTC = true
	// DefaultVerbose @todo
	DefaultVerbose = 0
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

// Dupless @todo
type Dupless struct {
	cache     string
	chunk     uint
	separator string
	dirReport bool
	exclude   string
	extra     bool
	freq      uint
	hash      string
	help      bool
	iexclude  string
	// rename to include
	mask         string
	minDirLength uint
	minFiles     uint
	minSize      uint64
	sizeReport   bool
	recursive    bool
	utc          bool
	verbose      uint

	cacheFH  *os.File
	comma    rune
	excludes []string
	lastDev  string
	masks    []string
	p        *message.Printer
	path     string
	dev      string

	// device stats:
	hits        uint
	skipped     uint
	directories uint
	matched     uint
	errors      uint
	ignored     uint
	dirs        map[string]*Dir
	errorDirs   map[string][]*ErrorRec
	ignoredDirs map[string][]*IgnoredRec
	//			files[path] = *file.File
	files map[string]*file.File
	//          uniques[uniqueID] = paths[]
	uniques map[string][]string
	//          sizes[size][uniqueIDs] = paths[]
	sizes map[uint64]map[string][]*file.File
	//          hashes[size][hash] = *file.File[]
	hashes map[uint64]map[string][]*file.File
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
	minDirLength := DefaultMinDirLength
	if runtime.GOOS == "windows" {
		// c:/
		minDirLength = 3
	}

	flag.StringVar(&d.cache, "cache", DefaultCache, "Cache filename")
	flag.UintVar(&d.chunk, "chunk", DefaultChunk, "Hash chunk")
	flag.BoolVar(&d.dirReport, "dir_report", DefaultDirReport, "Report by directory")
	flag.StringVar(&d.exclude, "exclude", DefaultExclude, "Regexs of Directories/files to exclude, separated by |")
	flag.StringVar(&d.iexclude, "iexclude", DefaultIexclude, "Regexs of Directories/files to exclude, separated by |")
	flag.BoolVar(&d.extra, "extra", DefaultExtra, "Cache extra attributes")
	flag.UintVar(&d.freq, "frequency", DefaultFrequency, "Reporting frequency")
	flag.StringVar(&d.hash, "hash", DefaultHash, "Hash type")
	flag.BoolVar(&d.help, "help", false, "Display help")
	flag.StringVar(&d.mask, "mask", DefaultMask, "File mask")
	flag.UintVar(&d.minDirLength, "min_dir_len", minDirLength, "Minimum directory length")
	flag.UintVar(&d.minFiles, "min_files", DefaultMinFiles, "Minimum files")
	flag.Uint64Var(&d.minSize, "min_size", DefaultMinSize, "Minimum file size")
	flag.BoolVar(&d.recursive, "recursive", DefaultRecursive, "Report directories recursively")
	flag.BoolVar(&d.sizeReport, "size_report", DefaultSizeReport, "Report by size")
	flag.StringVar(&d.separator, "separator", DefaultSeparator, "Field separator")
	flag.BoolVar(&d.utc, "utc", DefaultUTC, "Report times in UTC")
	flag.UintVar(&d.verbose, "verbose", DefaultVerbose, "Be more verbose")

	flag.Parse()

	d.hash = strings.ToLower(d.hash)

	if runtime.GOOS != "windows" {
		d.excludes = []string{
			"^/dev$",
			"^/proc$",
			"^/run$",
			"^/sys$",
		}
	} else {
		d.excludes = []string{
			`(?i)^[A-Z]:/$Recycle\.bin`,
			`(?i)^[A-Z]:/System Volume Information`,
		}
	}

	if d.exclude != "" {
		a := strings.Split(d.exclude, "|")
		for _, s := range a {
			s = util.NormalizePath(s)
			d.excludes = append(d.excludes, s)
		}
	}

	if d.iexclude != "" {
		a := strings.Split(d.iexclude, "|")
		for _, s := range a {
			s = util.NormalizePath(s)
			d.excludes = append(d.excludes, "(?i)"+s)
		}
	}

	//util.Dump("d.excludes=", d.excludes)

	if d.mask != "" {
		a := strings.Split(d.mask, "|")
		for _, s := range a {
			d.masks = append(d.masks, s)
		}
	}

	value, _ /*multibyte*/, _ /*tail*/, err := strconv.UnquoteChar(d.separator, 0)
	if err != nil {
		panic(err)
	}
	d.comma = value

	d.dirs = make(map[string]*Dir)
	d.errorDirs = make(map[string][]*ErrorRec)
	d.ignoredDirs = make(map[string][]*IgnoredRec)
	d.p = message.NewPrinter(language.English)
	d.files = make(map[string]*file.File)
	d.uniques = make(map[string][]string)
	d.sizes = make(map[uint64]map[string][]*file.File)
}

func (d *Dupless) addPath(path string, size uint64) {
	for {
		dir := util.Dirname(path)
		if uint(len(dir)) < d.minDirLength {
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
		if !d.recursive {
			return
		}
		path = dir
	}
}

func (d *Dupless) progress(final bool) {
	if !final && (d.freq == 0 || d.hits%d.freq != 0) {
		return
	}

	d.p.Printf("\r%11d %11d %11d %11d %11d %s", d.skipped, d.matched, d.directories, d.ignored, d.errors, d.dev)

	if final {
		fmt.Println("")
	}
}

func (d *Dupless) addError(path string, s string) {
	dir := util.Dirname(path)
	errorRec := ErrorRec{Path: path, Error: s}
	d.errorDirs[dir] = append(d.errorDirs[dir], &errorRec)
	d.errors++
	if d.verbose > 0 {
		fmt.Fprintf(os.Stderr, "\n%s\n", s)
	}
}

func (d *Dupless) addIgnore(path string, typ string) {
	dir := util.Dirname(path)
	IgnoredRec := IgnoredRec{Path: path, Type: typ}
	d.ignoredDirs[dir] = append(d.ignoredDirs[dir], &IgnoredRec)
	d.ignored++
	if d.verbose > 0 {
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
	if len(d.ignoredDirs) == 0 {
		return
	}

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
				err := f.Read(uint64(d.chunk))
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
			if uint(len(files)) < d.minFiles {
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

	switch d.hash {
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
		fmt.Fprintf(os.Stderr, "\nUnknown hash format: '%s'\n", d.hash)
		os.Exit(1)
	}

	return nil
}

func (d *Dupless) getHashes() {
	//start := time.Now()
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

	//util.Dump("d.files=", d.files)
	//util.Dump("d.uniques=", d.uniques)

	for size, uniques := range d.sizes {
		if uint(len(uniques)) < d.minFiles {
			delete(d.sizes, size)
		}
	}
	//util.Dump("d.sizes=", d.sizes)

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

	if len(sizes) < 1 {
		return
	}

	loops := (sizes[0] / uint64(d.chunk)) + 1

	loop := 0
	read := uint64(0)
	for {
		loop++
		scanning, total := d.countFiles(hashes)
		d.p.Printf("Loop %d of %d: %d of %d bytes read: scanning %d of %d files (%d unique sizes)\n", loop, loops, read, sizes[0], scanning, total, len(hashes))
		read += uint64(d.chunk)
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
			fmt.Println("No files left to hash")
			break
		}
	}
	//elapsed := time.Since(start)
	//_, total := d.countFiles(hashes)
	//fmt.Printf("\nHashed %d files in %s\n", total, elapsed)
	//util.Pause()
	//util.Dump("hashes=", hashes)
	d.hashes = hashes
}

func (d *Dupless) visit(path string, fi os.FileInfo, err error) error {
	path = util.NormalizePath(path)
	if d.verbose > 0 {
		fmt.Printf("Opening %s\n", path)
	}
	if err != nil {
		if d.verbose > 0 {
			fmt.Fprintf(os.Stderr, "\nError on '%s': %s\n", path, err)
		}
	}

	err = nil
	for {
		d.path = path
		d.hits++

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
				d.skipped++
				break
			}
		}

		if fi != nil {
			if fi.IsDir() {
				d.directories++
				break
			}
			if fi.Mode()&os.ModeSymlink != 0 {
				d.ignored++
				break
			}
			if !fi.Mode().IsRegular() {
				d.addIgnore(path, fi.Mode().String())
				break
			}

			if uint64(fi.Size()) <= d.minSize {
				d.skipped++
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
			d.directories++
			break
		}
		if f.IsSymlink() {
			d.ignored++
			break
		}
		if !f.IsRegular() {
			d.addIgnore(path, f.Type())
			break
		}

		if f.Size() <= d.minSize {
			d.skipped++
			break
		}

		d.dev, _ = f.VolumeName()
		if d.lastDev != d.dev {
			if d.lastDev != "" {
				fmt.Printf("\nSkipping %s as it is on device %s\n", path, d.dev)
				return filepath.SkipDir
			}
			d.lastDev = d.dev
		}

		d.matched++
		d.files[path] = f
		break
	}

	d.progress(false)
	return err
}

func (d *Dupless) resetCounters() {
	d.skipped = 0
	d.directories = 0
	d.matched = 0
	d.errors = 0
	d.ignored = 0
}

func usage() {
	fmt.Fprintf(os.Stderr, "\nUsage: %s [options] path [path2] ...\nOptions:\n\n", os.Args[0])
	flag.PrintDefaults()
}

func main() {
	basename := filepath.Base(os.Args[0])
	progname := strings.TrimSuffix(basename, filepath.Ext(basename))

	fmt.Printf("%s: Version %s (%s)\n", progname, version.VERSION, version.GITCOMMIT)
	fmt.Printf("Built with %s for %s/%s (%d CPUs/%d GOMAXPROCS)\n",
		runtime.Version(),
		runtime.GOOS,
		runtime.GOARCH,
		runtime.NumCPU(),
		runtime.GOMAXPROCS(-1))

	var dupless Dupless
	dupless.init()
	flag.Parse()

	if len(flag.Args()) == 0 || dupless.help {
		usage()
		return
	}

	dupless.p.Printf("\nMinimum size: %d\n\n", dupless.minSize)

	fmt.Printf("    skipped     matched directories     ignored      errors device\n")
	fmt.Printf("----------- ----------- ----------- ----------- ----------- ------\n")

	start := time.Now()

	args := flag.Args()
	var newargs []string
	for _, arg := range args {
		if arg != "*" {
			newargs = append(newargs, arg)
		} else {
			volumes, err := file.GetVolumes()
			if err != nil {
				fmt.Fprintln(os.Stderr, "\nGetVolumes() returned:", err)
			}
			for _, volume := range volumes {
				newargs = append(newargs, volume)
			}
		}
	}

	for i, arg := range newargs {
		if i > 0 {
			fmt.Println("")
		}
		dupless.resetCounters()
		dupless.lastDev = ""
		if runtime.GOOS == "windows" {
			if len(arg) == 2 && arg[1] == ':' {
				arg += string(os.PathSeparator)
			}
		}
		err := filepath.Walk(arg, dupless.visit)
		if err != nil {
			fmt.Fprintln(os.Stderr, "\nWalk returned:", err)
		}
	}

	dupless.progress(true)

	elapsed := time.Since(start)

	if len(dupless.files) < 1 {
		fmt.Printf("No files found\n")
		os.Exit(0)
	}

	start2 := time.Now()

	dupless.getHashes()

	elapsed2 := time.Since(start2)
	elapsed3 := time.Since(start)

	if len(dupless.hashes) < 1 {
		fmt.Printf("No duplicate files found\n")
		os.Exit(0)
	}

	if dupless.dirReport {
		dupless.reportByDir()
	}
	if dupless.sizeReport {
		dupless.reportBySize()
	}
	fmt.Printf("\nFound %d matching files in %s\n", len(dupless.files), elapsed)
	_, total := dupless.countFiles(dupless.hashes)
	fmt.Printf("Hashed %d files in %s\n", total, elapsed2)
	fmt.Printf("Total elapsed time: %s\n", elapsed3)
}
