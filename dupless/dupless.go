// Package dupless contains the duplication analysis engine
package dupless

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
	"github.com/rasa/godupless/types"
	"github.com/rasa/godupless/util"
	"github.com/rasa/jibber_jabber"
	"golang.org/x/text/language"
	"golang.org/x/text/message"
)

const (
	// @todo move to Config?
	// from https://github.com/minio/highwayhash/blob/master/examples_test.go#L17
	hashKey = "000102030405060708090A0B0C0D0E0FF0E0D0C0B0A090807060504030201000"
	// DefaultLanguage is the default language to display numbers in
	DefaultLanguage = "en" // English
	// MinChunk minimum chunk size, should be a power of 2
	MinChunk = 2 //2 << 12 // 4K
	// MaxChunk maximum chunk size, should be a power of 2
	MaxChunk = 2 << 29 // 1G
	// MinMinFiles Mininum number of the minimum files to compare to hash
	MinMinFiles = 2
)

// Config Config settings for Dupless class
type Config struct {
	//Cache string
	// Chunk @todo
	Chunk uint
	//Separator string
	// Exclude @todo
	Exclude string
	//Extra bool
	// Freq @todo
	Freq uint
	// Hash @todo
	Hash string
	// Help @todo
	Help bool
	// Iexclude @todo
	Iexclude string
	// Mask @todo
	Mask string
	// @todo add an imask option?
	// MinDirLength @todo
	MinDirLength uint
	// MinFiles @todo
	MinFiles uint
	// MinSize @todo
	MinSize uint64
	// Recursive @todo
	Recursive bool
	// Utc bool
	// Verbose @todo
	Verbose int

	// DiffReport @todo
	DiffReport bool
	// DirReport @todo
	DirReport bool
	// ErrorReport @todo
	ErrorReport bool
	// HardLinkReport @todo
	HardLinkReport bool
	// IgnoredReport @todo
	IgnoredReport bool
	// SizeReport @todo
	SizeReport bool
}

// default (initial) config settings
var config = Config{
	// Cache: "godupless.cache",
	Chunk: 2 << 16, // 2<<16 = 2^17 = 131,072
	// Separator: ",",
	//Exclude: "",
	// Extra: false,
	Freq: 100,
	Hash: "xxhash",
	//Help: false,
	//Iexclude: "",
	//Mask: "",
	MinFiles: 2,
	MinSize:  2 << 20, // 2<<20 = 2^21 = 2,097,152
	//Recursive: false,
	// Utc: false,
	//Verbose: 0,

	//DiffReport: false,
	//DirReport: false,
	//ErrorReport: false,
	//HardLinkReport: false,
	//IgnoredReport: false,
	SizeReport: true,
}

// ErrorRec  Error report record (for scanning relared errors)
type ErrorRec struct {
	// Path contains the full path of the file that generated an error
	Path string
	// Error contains the error message related to the file
	Error string
}

// ErrorFiles Error report record (for hashing relared errors)
type ErrorFiles struct {
	// Files has the list of files of the same size as a file that errored
	Files []*file.File
	// Error has the error that caused this group of files to be skipped
	Err error
}

// IgnoredRec Ignored report record
type IgnoredRec struct {
	// Path contains the full path of the ignore file
	Path string
	// Type contains the type of ignored file (symlink, named pipe, etc)
	Type string
}

// FileStats @todo
type FileStats struct {
	// Start @todo
	Start time.Time
	// Duration @todo
	Duration time.Duration
	// @todo
	// volume stats:
	// Hits @todo
	Hits uint
	// Skips @todo
	Skips uint
	// Directories @todo
	Directories uint
	// Matches @todo
	Matches uint
	// Errors @todo
	Errors uint
	// Ignores @todo
	Ignores uint
}

// HashStats @todo
type HashStats struct {
	// "Loop %d of %d: %d of %d bytes read: scanning %d of %d files (%d unique sizes)\n",
	// Start @todo
	Start time.Time
	//  @todo
	Duration time.Duration
	// Left @todo
	Left time.Duration
	// Finish Estimated finish time
	Finish time.Time
	// MbPerSecond @todo
	MbPerSecond float64
	// Loop  current loop index
	Loop uint
	// Loops total number of loops
	Loops uint
	// ReadBytes bytes read of each file
	ReadBytes uint64
	// MaxBytes size of largest file to be hashed
	MaxBytes uint64
	// DupFiles number of duplicate files hashed so far
	DupFiles int
	// UniqueFiles number of unique files hashed so far
	UniqueFiles int
	// HashedFiles number of files hashed so far
	HashedFiles int
	// ErrorFiles number of failed that were not successfully hashed due to an error
	ErrorFiles int
	// RemainingFiles number of files remaining to be hashed
	RemainingFiles int
	// RemainingSizes  number of sizes remaining
	RemainingSizes int
	// TotalFiles number of files that were scanned
	TotalFiles int
	// TotalReadBytes total bytes read so far
	TotalReadBytes uint64
	// TotalRemainingBytes total bytes remainging to be read
	TotalRemainingBytes uint64

	// @todo ?
	// dupBytes uint64 // total bytes of duplicate files hashed so far
	// uniqueBytes uint64 // total bytes of unique files hashed so far
	// hashedSizes uint64 // number of sizes that have been hashed so far
	// totalSizes uint64
}

func highway64New() hash.Hash {
	h, _ := highwayhash.New64([]byte(hashKey))
	return h
}

func highway128New() hash.Hash {
	h, _ := highwayhash.New128([]byte(hashKey))
	return h
}

func highway256New() hash.Hash {
	h, _ := highwayhash.New([]byte(hashKey))
	return h
}

func xxhashNew() hash.Hash {
	return xxhash.New()
}

// Hashmap list of hash methods
var Hashmap = map[string]func() hash.Hash{
	"highway":    highway256New,
	"highway64":  highway64New,
	"highway128": highway128New,
	"highway256": highway256New,
	"md5":        md5.New,
	"sha1":       sha1.New,
	"sha256":     sha256.New,
	"sha512":     sha512.New,
	"xxhash":     xxhashNew,
}

// Dupless @todo
type Dupless struct {
	// Config @todo
	Config Config
	// Fstats @todo
	Fstats FileStats
	// Hstats @todo
	Hstats HashStats
	// P @todo
	P *message.Printer
	// Args @todo
	Args []string
	// Excludes @todo
	Excludes []string
	// Masks @todo
	Masks []string
	// Path @todo
	Path string
	// Dev @todo
	Dev string
	// LastDev @todo
	LastDev string
	// Volume @todo
	Volume string
	// HashFunc @todo
	HashFunc func() hash.Hash
	//CacheFH *os.File
	// Comma   rune

	// ErrorDirs @todo
	// ErrorDirs[dir] = []*ErrorRec{Path, Error}
	ErrorDirs map[string][]*ErrorRec
	// IgnoredDirs @todo
	// IgnoredDirs[dir] = []*IgnoredRec{Path, Type}
	IgnoredDirs map[string][]*IgnoredRec
	// Files @todo
	// Files[path] = *file.File
	Files map[string]*file.File
	// Uniques @todo
	// Uniques[uniqueID] = []paths
	Uniques map[string][]string
	// Sizes @todo
	// Sizes[size][uniqueIDs] = *file.File
	Sizes map[uint64]map[string]*file.File
	// Hashes @todo
	// Hashes[size][hash] = []*file.File
	Hashes map[uint64]map[string][]*file.File
	// Dups @todo
	// Dups[size][hash] = []*file.File
	Dups map[uint64]map[string][]*file.File
	// Errors @todo
	// Errors[path]error
	Errors map[uint64]map[string]*ErrorFiles
}

// Init @todo
func (d *Dupless) Init() {
	d.ErrorDirs = make(map[string][]*ErrorRec)
	d.IgnoredDirs = make(map[string][]*IgnoredRec)
	d.Files = make(map[string]*file.File)
	d.Uniques = make(map[string][]string)
	d.Sizes = make(map[uint64]map[string]*file.File)
	d.Hashes = make(map[uint64]map[string][]*file.File)
	d.Dups = make(map[uint64]map[string][]*file.File)
	d.Errors = make(map[uint64]map[string]*ErrorFiles)

	d.Config = config

	_, err := hex.DecodeString(hashKey)
	if err != nil {
		fmt.Printf("Invalid hash key %s: %s\n", hashKey, err)
		os.Exit(1)
	}

	userLanguage, err := jibber_jabber.DetectLanguage()
	if err != nil {
		userLanguage = DefaultLanguage
	}
	//fmt.Println("Language:", userLanguage)

	tagLanguage := language.Make(userLanguage)
	d.P = message.NewPrinter(tagLanguage) // language.English)

	chunk := fmt.Sprintf("Hash chunk size (%d to %d)", MinChunk, MaxChunk)

	var hashes []string
	for k := range Hashmap {
		hashes = append(hashes, k)
	}
	hash := "Hash type: " + strings.Join(hashes, ",")

	// flag.StringVar(&d.Config.Cache, "cache", d.Config.Cache, "Cache filename")
	flag.UintVar(&d.Config.Chunk, "chunk", d.Config.Chunk, chunk)
	flag.StringVar(&d.Config.Exclude, "exclude", d.Config.Exclude, "Regex(s) of directories/files to exclude, separated by |")
	flag.StringVar(&d.Config.Iexclude, "iexclude", d.Config.Iexclude, "Regex(s) of directories/files to exclude, separated by |")
	//flag.BoolVar(&d.Config.Extra, "extra", d.Config.Extra, "Cache extra attributes")
	flag.UintVar(&d.Config.Freq, "frequency", d.Config.Freq, "Reporting frequency")
	flag.StringVar(&d.Config.Hash, "hash", d.Config.Hash, hash)
	flag.BoolVar(&d.Config.Help, "help", d.Config.Help, "Display help")
	flag.StringVar(&d.Config.Mask, "mask", d.Config.Mask, "File mask(s), seperated by |")
	flag.UintVar(&d.Config.MinFiles, "min_files", d.Config.MinFiles, "Minimum files to compare")
	flag.Uint64Var(&d.Config.MinSize, "min_size", d.Config.MinSize, "Minimum file size")
	flag.BoolVar(&d.Config.Recursive, "recursive", d.Config.Recursive, "Report directories recursively")
	//flag.StringVar(&d.Config.Seperator, "separator", d.Config.Seperator, "Field separator")
	// flag.BoolVar(&d.Config.Utc, "utc", d.Config.Utc, "Report times in UTC")
	flag.IntVar(&d.Config.Verbose, "verbose", d.Config.Verbose, "Increase log verbosity")

	flag.BoolVar(&d.Config.DiffReport, "diff_report", d.Config.DiffReport, "Report on differences between directories containing duplicate files")
	flag.BoolVar(&d.Config.DirReport, "dir_report", d.Config.DirReport, "Report summary of duplicates by directory")
	flag.BoolVar(&d.Config.ErrorReport, "error_report", d.Config.ErrorReport, "Report of errors")
	flag.BoolVar(&d.Config.HardLinkReport, "hard_link_report", d.Config.HardLinkReport, "Report on hard link differences")
	flag.BoolVar(&d.Config.IgnoredReport, "ignored_report", d.Config.IgnoredReport, "Report of ignored files")
	flag.BoolVar(&d.Config.SizeReport, "size_report", d.Config.SizeReport, "Report duplicates by size")

	flag.Parse()

	if d.Config.Chunk < MinChunk || d.Config.Chunk > MaxChunk {
		fmt.Printf("Chunk must be between %d and %d", MinChunk, MaxChunk)
		os.Exit(1)
	}

	d.Config.Hash = strings.ToLower(d.Config.Hash)
	_, ok := Hashmap[d.Config.Hash]
	if !ok {
		fmt.Println("Hash must be one of ", hash)
		os.Exit(1)
	}

	if d.Config.MinFiles < MinMinFiles {
		fmt.Printf("Mininum files to compare must be %d or greater", MinMinFiles)
		os.Exit(1)
	}

	// @todo ignore all hidden/system directories?
	d.Excludes = file.ExcludePaths

	if d.Config.Exclude != "" {
		a := strings.Split(d.Config.Exclude, "|")
		for _, s := range a {
			s = util.NormalizePath(s)
			d.Excludes = append(d.Excludes, s)
		}
	}

	if d.Config.Iexclude != "" {
		a := strings.Split(d.Config.Iexclude, "|")
		for _, s := range a {
			s = util.NormalizePath(s)
			d.Excludes = append(d.Excludes, "(?i)"+s)
		}
	}

	//util.Dump("d.Excludes=", d.Excludes)

	if d.Config.Mask != "" {
		a := strings.Split(d.Config.Mask, "|")
		for _, s := range a {
			d.Masks = append(d.Masks, s)
		}
	}

	// value, _ /*multibyte*/, _ /*tail*/, err := strconv.UnquoteChar(d.Config.Separator, 0)
	/*
		if err != nil {
			panic(err)
		}
		d.Comma = value
	*/
}

// AddError @todo
func (d *Dupless) AddError(path string, s string) {
	dir := util.Dirname(path)
	errorRec := ErrorRec{Path: path, Error: s}
	d.ErrorDirs[dir] = append(d.ErrorDirs[dir], &errorRec)
	d.Fstats.Errors++
	if d.Config.Verbose > 0 {
		fmt.Fprintln(os.Stderr, "\n", s)
	}
}

// AddIgnore @todo
func (d *Dupless) AddIgnore(path string, typ string) {
	dir := util.Dirname(path)
	IgnoredRec := IgnoredRec{Path: path, Type: typ}
	d.IgnoredDirs[dir] = append(d.IgnoredDirs[dir], &IgnoredRec)
	d.Fstats.Ignores++
	if d.Config.Verbose > 0 {
		fmt.Fprintf(os.Stderr, "\nSkipping %q: %s\n", path, typ)
	}
}

// AddDups @todo
func (d *Dupless) AddDups(files []*file.File) {
	hash := files[0].Hash()
	if hash == "" {
		return
	}
	for _, f := range files[1:] {
		if f.Hash() != hash {
			return
		}
	}
	size := files[0].Size()
	_, ok := d.Dups[size]
	if !ok {
		d.Dups[size] = make(map[string][]*file.File)
	}
	d.Dups[size][hash] = files
}

// ReadSize @todo
func (d *Dupless) ReadSize(size uint64) {
	var err error
	hashmap := d.Hashes[size]

Loop:
	for {
		for _, files := range hashmap {
			for _, f := range files {
				//fmt.Printf("Opening %s\n", f.Path())
				err = f.Open()
				if err != nil {
					break Loop
				}
			}
		}

		var read uint64
		for read < size {
			read += uint64(d.Config.Chunk)
			finished := true
			for _, files := range hashmap {
				eof := true
				for _, f := range files {
					if !f.Opened() {
						continue
					}
					//fmt.Printf("Reading %d bytes from %s\n", d.Chunk, f.Path())
					err = f.Read(uint64(d.Config.Chunk))
					d.Hstats.TotalReadBytes += uint64(d.Config.Chunk)
					if err == io.EOF {
						f.Close()
						err = nil
						continue
					}
					if err != nil {
						// don't set eof to false on errors
						break Loop
					}
					eof = false
				}
				if !eof {
					finished = false
					continue
				}
			}
			if finished {
				break
			}
			hashes := make(map[string][]*file.File)
			for _, files := range hashmap {
				for _, f := range files {
					hash := f.Hash()
					_, ok := hashes[hash]
					if !ok {
						hashes[hash] = make([]*file.File, 0)
					}
					hashes[hash] = append(hashes[hash], f)
				}
			}
			finished = true
			for _, files := range hashes {
				if len(files) > 1 {
					finished = false
					break
				}
				files[0].Close()
				files[0].ResetHash()
			}
			if finished {
				break
			}
		}

		break
	}

	if err == nil {
		for _, files := range hashmap {
			d.AddDups(files)
		}
	} else {
		affected := 0
		d.Errors[size] = make(map[string]*ErrorFiles)
		for hash, files := range hashmap {
			affected += len(files)
			d.Errors[size][hash] = &ErrorFiles{Files: files, Err: err}
		}
		fmt.Printf("\nCannot hash %d files: %s\n", affected, err)
	}

	for _, files := range hashmap {
		for _, f := range files {
			f.Close()
			f.ResetHash()
		}
	}

	for hash := range hashmap {
		delete(d.Hashes[size], hash)
	}
	delete(d.Hashes, size)
}

// CalculateStats @todo
func (d *Dupless) CalculateStats() {
	d.Hstats.RemainingSizes = len(d.Hashes)

	d.Hstats.ErrorFiles = 0
	for _, hashmap := range d.Errors {
		for _, errorFiles := range hashmap {
			d.Hstats.ErrorFiles += len(errorFiles.Files)
		}
	}

	d.Hstats.DupFiles = 0
	for _, hashmap := range d.Dups {
		for _, files := range hashmap {
			d.Hstats.DupFiles += len(files)
		}
	}

	d.Hstats.Loops = d.Hstats.Loop + uint(len(d.Hashes))

	d.Hstats.Duration = time.Since(d.Hstats.Start)
	seconds := d.Hstats.Duration.Seconds()
	if seconds > 0.0 {
		bytesPerSecond := float64(d.Hstats.TotalReadBytes) / seconds
		d.Hstats.MbPerSecond = bytesPerSecond / 1000000
		remainingSeconds := int64(float64(d.Hstats.TotalRemainingBytes) / bytesPerSecond)
		d.Hstats.Left = time.Duration(remainingSeconds) * time.Second
		d.Hstats.Finish = time.Now().Add(d.Hstats.Left)
	}

	d.Hstats.RemainingFiles = 0

	if len(d.Hashes) > 0 {
		d.Hstats.MaxBytes = 0
		d.Hstats.TotalRemainingBytes = 0
		for size, hashmap := range d.Hashes {
			if size > d.Hstats.MaxBytes {
				d.Hstats.MaxBytes = size
			}

			for _, files := range hashmap {
				d.Hstats.RemainingFiles += len(files)
				d.Hstats.TotalRemainingBytes += size * uint64(len(files))
			}
		}
	}

	d.Hstats.HashedFiles = d.Hstats.TotalFiles - d.Hstats.RemainingFiles
	d.Hstats.UniqueFiles = d.Hstats.HashedFiles - d.Hstats.DupFiles - d.Hstats.ErrorFiles

	// @todo remove
	//util.Dump("\nhstats=", d.Hstats)
}

// ReadFiles @todo
func (d *Dupless) ReadFiles() {
	d.Hstats.Start = time.Now()

	i := 0
	sizes := make([]uint64, len(d.Hashes))
	for size := range d.Hashes {
		sizes[i] = size
		i++
	}
	sort.Sort(types.Uint64Slice(sizes))

	for _, size := range sizes {
		d.Hstats.Loop++
		d.Hstats.ReadBytes = size
		d.ReadSize(size)
		d.CalculateStats()
		if d.Config.Verbose >= 0 {
			d.P.Printf("\r%6d %6d %13d %13d %13d %13d %8s %8.2f", d.Hstats.Loop, d.Hstats.Loops, d.Hstats.ReadBytes, d.Hstats.MaxBytes, d.Hstats.RemainingFiles, d.Hstats.HashedFiles, d.Hstats.Left, d.Hstats.MbPerSecond)
			fmt.Println("")
		}
	}
	d.Hstats.Duration = time.Since(d.Hstats.Start)
}

// LoadHashmap @todo
func (d *Dupless) LoadHashmap() bool {
	for path, f := range d.Files {
		uniqueID := f.UniqueID()
		_, ok := d.Uniques[uniqueID]
		if !ok {
			d.Uniques[uniqueID] = make([]string, 0)
		}
		d.Uniques[uniqueID] = append(d.Uniques[uniqueID], path)
		_, ok = d.Sizes[f.Size()]
		if !ok {
			d.Sizes[f.Size()] = make(map[string]*file.File)
		}
		d.Sizes[f.Size()][uniqueID] = f
	}

	//util.Dump("d.Files=", d.Files)
	//util.Dump("d.Uniques=", d.Uniques)

	for size, uniques := range d.Sizes {
		if uint(len(uniques)) < d.Config.MinFiles {
			delete(d.Sizes, size)
		}
	}
	//util.Dump("d.Sizes=", d.Sizes)

	if len(d.Sizes) < 0 {
		return false
	}

	f := Hashmap[d.Config.Hash]
	h := f()
	defaultHash := fmt.Sprintf("%x", h.Sum(nil))

	for size, uniqueMap := range d.Sizes {
		d.Hashes[size] = make(map[string][]*file.File)
		d.Hashes[size][defaultHash] = make([]*file.File, 0)
		for _, f := range uniqueMap {
			d.Hashes[size][defaultHash] = append(d.Hashes[size][defaultHash], f)
		}
		d.Hstats.TotalFiles += len(uniqueMap)
		d.Hstats.TotalRemainingBytes += size * uint64(len(uniqueMap))
	}

	return true
}

// GetHashes @todo
func (d *Dupless) GetHashes() bool {
	if !d.LoadHashmap() {
		return false
	}

	//util.Dump("d.Hashes=", d.Hashes)

	fmt.Printf("Started:       %s\n\n", time.Now().Format("15:04:05"))

	fmt.Println("        Total       Current       Largest     Remaining        Hashed     Time      MB/")
	fmt.Println("  Pass Passes          Size          Size         Files         Files     Left   Second")
	fmt.Println("------ ------ ------------- ------------- ------------- ------------- -------- --------")
	//           123,56 123,56 3,109,765,321 3,109,765,321 3,109,765,321 3,109,765,321 12345678 12345.78

	d.ReadFiles()
	return len(d.Dups) > 0
}

// Progress @todo
func (d *Dupless) Progress(force bool) {
	if !force && (d.Config.Freq == 0 || d.Fstats.Hits%d.Config.Freq != 0) {
		return
	}

	dev := d.Dev
	if d.Volume > "" && dev != d.Volume {
		dev += " (" + d.Volume + ")"
	}
	d.P.Printf("\r%11d %11d %11d %11d %11d %s", d.Fstats.Skips, d.Fstats.Matches, d.Fstats.Directories, d.Fstats.Ignores, d.Fstats.Errors, dev)
}

// Visit @todo
func (d *Dupless) Visit(path string, fi os.FileInfo, err error) error {
	path = util.NormalizePath(path)
	if d.Config.Verbose > 0 {
		fmt.Printf("Opening %q\n", path)
	}
	if err != nil {
		if d.Config.Verbose > 0 {
			fmt.Fprintf(os.Stderr, "\nError on %q: %s\n", path, err)
		}
	}

	hfunc := Hashmap[d.Config.Hash]

	err = nil
	for {
		d.Path = path
		d.Fstats.Hits++

		for _, exclude := range d.Excludes {
			ok, e := regexp.MatchString(exclude, path)
			if e != nil {
				s := fmt.Sprintf("Failed to match %q via %q: %s", path, exclude, e)
				d.AddError(path, s)
			}
			if ok {
				d.AddIgnore(path, "excluded")
				return filepath.SkipDir
			}
		}

		if len(d.Masks) > 0 {
			//_, file := filepath.Split(path)
			matched := false
			for _, mask := range d.Masks {
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
				d.Fstats.Skips++
				break
			}
		}

		if fi != nil {
			if fi.IsDir() {
				d.Fstats.Directories++
				break
			}
			if fi.Mode()&os.ModeSymlink != 0 {
				d.Fstats.Ignores++
				break
			}
			if !fi.Mode().IsRegular() {
				d.AddIgnore(path, fi.Mode().String())
				break
			}

			if uint64(fi.Size()) <= d.Config.MinSize {
				d.Fstats.Skips++
				break
			}
		}

		f, e := file.NewFile(path, fi, hfunc())
		if e != nil {
			s := fmt.Sprintf("Cannot stat %q: %s", path, e)
			d.AddError(path, s)
			break
		}

		if f.IsDir() {
			d.Fstats.Directories++
			break
		}
		if f.IsSymlink() {
			d.Fstats.Ignores++
			break
		}
		if !f.IsRegular() {
			d.AddIgnore(path, f.Type())
			break
		}

		if f.Size() <= d.Config.MinSize {
			d.Fstats.Skips++
			break
		}

		d.Dev, _ = f.VolumeName()
		if d.LastDev != d.Dev {
			// @todo add option to include/exclude cross-device files
			if d.LastDev != "" {
				// @todo log skipped files
				fmt.Printf("\nSkipping %q as it is on device %s\n", path, d.Dev)
				return filepath.SkipDir
			}
			d.LastDev = d.Dev
		}

		d.Fstats.Matches++
		d.Files[path] = f
		break
	}

	d.Progress(false)
	return err
}

// ResetCounters @todo
func (d *Dupless) ResetCounters() {
	d.Fstats.Skips = 0
	d.Fstats.Directories = 0
	d.Fstats.Matches = 0
	d.Fstats.Errors = 0
	d.Fstats.Ignores = 0
	d.LastDev = ""
}

// Header @todo
func (d *Dupless) Header() {
	fmt.Println("")
	fmt.Printf("Arguments:     %s\n", strings.Join(d.Args, ","))
	d.P.Printf("Chunk size:    %d\n", d.Config.Chunk)
	fmt.Printf("Hash format:   %s\n", d.Config.Hash)
	d.P.Printf("Minimum files: %d\n", d.Config.MinFiles)
	d.P.Printf("Minimum size:  %d\n", d.Config.MinSize)
	fmt.Printf("Masks:         %v\n", d.Config.Mask)
	fmt.Printf("Recursive:     %v\n", d.Config.Recursive)
	fmt.Printf("Verbosity:     %d\n", d.Config.Verbose)
	fmt.Println("")

	fmt.Printf("Started:       %s\n\n", time.Now().Format("15:04:05"))

	fmt.Println("    Skipped     Matched Directories     Ignored      Errors Device")
	fmt.Println("----------- ----------- ----------- ----------- ----------- ------")
}

// ProcessArgs @todo
func (d *Dupless) ProcessArgs() bool {
	fmt.Println("Scanning volumes...")
	volumes, err := file.GetVolumes()
	if err != nil {
		fmt.Fprintln(os.Stderr, "Failed to get volumes: ", err)
	}
	for _, arg := range flag.Args() {
		if arg == "*" {
			for _, volume := range volumes {
				d.Args = append(d.Args, volume)
			}
			continue
		}
		if runtime.GOOS == "windows" {
			if len(arg) == 2 && arg[1] == ':' {
				arg += string(os.PathSeparator)
			}
		}
		arg = filepath.Clean(arg)
		d.Args = append(d.Args, arg)
	}
	return true
}

// FindFiles @todo
func (d *Dupless) FindFiles() bool {
	d.Fstats.Start = time.Now()
	for i, arg := range d.Args {
		if i > 0 {
			fmt.Println("")
		}
		d.ResetCounters()
		if file.IsVolume(arg) {
			d.Volume = arg
		} else {
			d.Volume = ""
		}
		err := filepath.Walk(arg, d.Visit)
		if err != nil {
			fmt.Fprintln(os.Stderr, "\nWalk returned: ", err)
		}
	}

	d.Progress(true)

	d.Fstats.Duration = time.Since(d.Fstats.Start)
	fmt.Printf("\nFound %d matching files in %s\n", len(d.Files), d.Fstats.Duration)
	return len(d.Files) > 0
}

// Footer @todo
func (d *Dupless) Footer() {
	d.CalculateStats()
	fmt.Printf("\nFound %d matching files in %s\n", len(d.Files), d.Fstats.Duration)
	fmt.Printf("Found %d duplicate files in %s\n", d.Hstats.HashedFiles, d.Hstats.Duration)
	fmt.Printf("Total elapsed time: %s\n", d.Fstats.Duration+d.Hstats.Duration)
}

// Run @godo
func (d *Dupless) Run() bool {
	d.Init()

	if len(flag.Args()) == 0 || d.Config.Help {
		d.Config.Help = true
		return false
	}

	d.ProcessArgs()

	d.Header()

	if !d.FindFiles() {
		fmt.Println("No matching files found")
		return true
	}

	if !d.GetHashes() {
		fmt.Println("No duplicate files found")
		return false
	}

	// d.Footer()
	return true
}
