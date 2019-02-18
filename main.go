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
	"github.com/rasa/jibber_jabber"
	"golang.org/x/text/language"
	"golang.org/x/text/message"
)

const (
	// DefaultLanguage is the default language to display numbers in
	DefaultLanguage = "en" // English
	// MinChunk minimum chunk size, should be a power of 2
	MinChunk = 2 //2 << 12 // 4K
	// MaxChunk maximum chunk size, should be a power of 2
	MaxChunk = 2 << 29 // 1G
	// MinMinFiles Mininum number of the minimum files to compare to hash
	MinMinFiles = 2
)

// Dir @todo
type Dir struct {
	// Count number of duplicate files in the directory
	Count uint64
	// Size total size of the duplicate files in the directory
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

// ErrorFiles @todo
type ErrorFiles struct {
	// Files has the list of files of the same size as a file that errored
	Files []*file.File
	// Error has the error that caused this group of files to be skipped
	Err error
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
	// Chunk @todo
	Chunk uint
	//separator string
	// Exclude @todo
	Exclude string
	//extra         bool
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
	// @todo add and imask option?
	// MinDirLength @todo
	MinDirLength uint
	// MinFiles @todo
	MinFiles uint
	// MinSize @todo
	MinSize uint64
	// Recursive @todo
	Recursive bool
	// utc          bool
	// Verbose @todo
	Verbose int

	// DirReport @todo
	DirReport bool
	// ErrorReport @todo
	ErrorReport bool
	// IgnoredReport @todo
	IgnoredReport bool
	// SizeReport @todo
	SizeReport bool
}

var config = Config{
	// Cache: "godupless.cache",
	Chunk: 2 << 16, // 2<<16 = 2^17 = 131,072
	// Separator: ",",
	//Exclude: "",
	// Extra: false,
	Freq: 100,
	Hash: "highway",
	//Help: false,
	//Iexclude: "",
	//Mask: "",
	MinFiles: 2,
	MinSize:  2 << 20, // 2<<20 = 2^21 = 2,097,152
	//Recursive: false,
	// Utc: false,
	//Verbose: 0,

	//DirReport: false,
	//ErrorReport: false,
	//IgnoredReport: false,
	SizeReport: true,
}

const hashKey = "0000000000000000000000000000000000000000000000000000000000000000"

var hashBytes []byte

func init() {
	hashBytes, _ = hex.DecodeString(hashKey)
}

func highway64New() hash.Hash {
	h, _ := highwayhash.New64(hashBytes)
	return h
}

func highway128New() hash.Hash {
	h, _ := highwayhash.New128(hashBytes)
	return h
}

func highway256New() hash.Hash {
	h, _ := highwayhash.New(hashBytes)
	return h
}

func xxhashNew() hash.Hash {
	return xxhash.New()
}

// FileStats @todo
type FileStats struct {
	// Start @todo
	Start time.Time
	// Duration @todo
	Duration time.Duration
	// // @todo
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
	// Finish @todo
	Finish time.Time
	// MbPerSecond @todo
	MbPerSecond float64
	// Loop @todo
	Loop uint // current loop index
	// Loops @todo
	Loops uint // total loops
	// ReadBytes @todo
	ReadBytes uint64 // bytes read of each file
	// MaxBytes @todo
	MaxBytes uint64 // size of largest file to be hashed
	// DupFiles @todo
	DupFiles int // number of duplicate files hashed so far
	// UniqueFiles @todo
	UniqueFiles int // number of unique files hashed so far
	// HashedFiles @todo
	HashedFiles int // number of unique files hashed so far
	// ErrorFiles @todo
	ErrorFiles int
	// RemainingFiles @todo
	RemainingFiles int // number of files remaining to be hashed
	// RemainingSizes @todo
	RemainingSizes int // number of sizes remaining
	// TotalFiles @todo
	TotalFiles int
	// TotalReadBytes @todo
	TotalReadBytes uint64 // total bytes read
	// TotalRemainingBytes @todo
	TotalRemainingBytes uint64 // total bytes remainging to be read

	// @todo:
	//dupBytes uint64 // total bytes of duplicate files hashed so far
	//uniqueBytes uint64 // total bytes of unique files hashed so far
	//hashedSizes uint64 // number of sizes that have been hashed so far
	// totalSizes uint64
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

	// Dirs[dir] = *Dir{Count, Size}
	// ErrorDirs[dir] = []*ErrorRec{Path, Error}
	// IgnoredDirs[dir] = []*IgnoredRec{Path, Type}
	// Files[path] = *file.File
	// Uniques[uniqueID] = []paths
	// Sizes[size][uniqueIDs] = *file.File
	// Hashes[size][hash] = []*file.File
	// Dups[size][hash] = []*file.File
	// Errors[path]error

	// Dirs @todo
	Dirs map[string]*Dir
	// ErrorDirs @todo
	ErrorDirs map[string][]*ErrorRec
	// IgnoredDirs @todo
	IgnoredDirs map[string][]*IgnoredRec
	// Files @todo
	Files map[string]*file.File
	// Uniques @todo
	Uniques map[string][]string
	// Sizes @todo
	Sizes map[uint64]map[string]*file.File
	// Hashes @todo
	Hashes map[uint64]map[string][]*file.File
	// Dups @todo
	Dups map[uint64]map[string][]*file.File
	// Errors @todo
	Errors map[uint64]map[string]*ErrorFiles
}

// Uint64Slice @todo
type Uint64Slice []uint64

func (p Uint64Slice) Len() int           { return len(p) }
func (p Uint64Slice) Less(i, j int) bool { return p[i] < p[j] }
func (p Uint64Slice) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }

// Sort is a convenience method
func (p Uint64Slice) Sort() { sort.Sort(p) }

// Init @todo
func (d *Dupless) Init() {
	d.Dirs = make(map[string]*Dir)
	d.ErrorDirs = make(map[string][]*ErrorRec)
	d.IgnoredDirs = make(map[string][]*IgnoredRec)
	d.Files = make(map[string]*file.File)
	d.Uniques = make(map[string][]string)
	d.Sizes = make(map[uint64]map[string]*file.File)
	d.Hashes = make(map[uint64]map[string][]*file.File)
	d.Dups = make(map[uint64]map[string][]*file.File)
	d.Errors = make(map[uint64]map[string]*ErrorFiles)

	d.Config = config

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
	flag.BoolVar(&d.Config.DirReport, "dir_report", d.Config.DirReport, "Report by directory")
	flag.BoolVar(&d.Config.ErrorReport, "error_report", d.Config.ErrorReport, "Report of errors")
	flag.StringVar(&d.Config.Exclude, "exclude", d.Config.Exclude, "Regex(s) of directories/files to exclude, separated by |")
	flag.StringVar(&d.Config.Iexclude, "iexclude", d.Config.Iexclude, "Regex(s) of directories/files to exclude, separated by |")
	//flag.BoolVar(&d.Config.Extra, "extra", d.Config.Extra, "Cache extra attributes")
	flag.UintVar(&d.Config.Freq, "frequency", d.Config.Freq, "Reporting frequency")
	flag.StringVar(&d.Config.Hash, "hash", d.Config.Hash, hash)
	flag.BoolVar(&d.Config.Help, "help", d.Config.Help, "Display help")
	flag.BoolVar(&d.Config.IgnoredReport, "ignored_report", d.Config.IgnoredReport, "Report of ignored files")
	flag.StringVar(&d.Config.Mask, "mask", d.Config.Mask, "File mask(s), seperated by |")
	flag.UintVar(&d.Config.MinFiles, "min_files", d.Config.MinFiles, "Minimum files to compare")
	flag.Uint64Var(&d.Config.MinSize, "min_size", d.Config.MinSize, "Minimum file size")
	flag.BoolVar(&d.Config.Recursive, "recursive", d.Config.Recursive, "Report directories recursively")
	flag.BoolVar(&d.Config.SizeReport, "size_report", d.Config.SizeReport, "Report by size")
	//flag.StringVar(&d.Config.Seperator, "separator", d.Config.Seperator, "Field separator")
	// flag.BoolVar(&d.Config.Utc, "utc", d.Config.Utc, "Report times in UTC")
	flag.IntVar(&d.Config.Verbose, "verbose", d.Config.Verbose, "Increase log verbosity")

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

// AddPath @todo
func (d *Dupless) AddPath(path string, size uint64) {
	for {
		dir := util.Dirname(path)
		if uint(len(dir)) < file.MinDirLength {
			return
		}
		if dir == path {
			return
		}
		_, ok := d.Dirs[dir]
		if !ok {
			d.Dirs[dir] = &Dir{Count: 1, Size: size}
		} else {
			d.Dirs[dir].Count++
			d.Dirs[dir].Size += size
		}
		if !d.Config.Recursive {
			return
		}
		path = dir
	}
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

// ReportByDir @todo
func (d *Dupless) ReportByDir() {
	fmt.Print("\nDuplication Report By Size/Directory\n\n")

	for size, hashmap := range d.Dups {
		for _, files := range hashmap {
			for _, f := range files {
				d.AddPath(f.Path(), size)
			}
		}
	}

	sizemap := make(map[uint64][]*DirCount)

	for dir, d := range d.Dirs {
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
				d.P.Printf("size: %d\n", size)
				emitSize = false
			}
			d.P.Printf("  %d: %v (%d files)\n", i+1, dircount.Dir, dircount.Count)
			files += dircount.Count
		}
	}
	d.P.Printf("\n%d files totaling %d bytes (%d bytes per file average)\n", files, totalSize, totalSize/files)
}

// ReportBySize @todo
func (d *Dupless) ReportBySize() {
	fmt.Print("\nDuplication Report By Size/Paths\n\n")

	i := 0
	sizes := make([]uint64, len(d.Dups))
	for size := range d.Dups {
		sizes[i] = size
		i++
	}
	sort.Sort(sort.Reverse(Uint64Slice(sizes)))

	var totalSize uint64
	var totalFiles uint64

	for _, size := range sizes {
		for hash, files := range d.Dups[size] {
			totalFiles += uint64(len(files))
			d.P.Printf("Size: %d (%s)\n", size, hash)
			for i, f := range files {
				totalSize += size
				var s string
				if f.Links() > 1 {
					s = fmt.Sprintf(" (%d hard links)", f.Links())
				}
				fmt.Printf("  %d: %q%s\n", i+1, f.Path(), s)
			}
		}
	}

	var avg uint64
	if totalFiles > 0 {
		avg = totalSize / totalFiles
	}

	d.P.Printf("\n%d files totaling %d bytes (%d bytes per file average)\n", totalFiles, totalSize, avg)
}

// ReportIgnored @todo
func (d *Dupless) ReportIgnored() {
	fmt.Print("\nIgnored Files/Directories Report\n\n")

	i := 0
	dirs := make([]string, len(d.IgnoredDirs))
	for dir := range d.IgnoredDirs {
		dirs[i] = dir
		i++
	}
	sort.Strings(dirs)

	for _, dir := range dirs {
		for _, other := range d.IgnoredDirs[dir] {
			fmt.Printf("%-10s: %q\n", other.Type, other.Path)
		}
	}
}

// ReportErrors @todo
func (d *Dupless) ReportErrors() {
	fmt.Print("\nError Files/Directories Report\n\n")

	i := 0
	dirs := make([]string, len(d.ErrorDirs))
	for dir := range d.ErrorDirs {
		dirs[i] = dir
		i++
	}
	sort.Strings(dirs)

	for _, dir := range dirs {
		for _, other := range d.ErrorDirs[dir] {
			fmt.Println(other.Error)
		}
	}

	fmt.Println("")

	for size, hashmap := range d.Errors {
		for hash, errorFiles := range hashmap {
			for _, f := range errorFiles.Files {
				fmt.Printf("%s: (size: %d, hash: %s)\n", f.Path(), size, hash)
			}
		}
	}
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

// ReadFiles @todo
func (d *Dupless) ReadFiles() (continueReading bool) {
	if len(d.Hashes) < 1 {
		return false
	}

	i := 0
	sizes := make([]uint64, len(d.Hashes))
	for size := range d.Hashes {
		sizes[i] = size
		i++
	}
	sort.Sort(Uint64Slice(sizes))

	for _, size := range sizes {
		d.Hstats.Loop++
		d.ReadSize(size)
		d.Hstats.ReadBytes = size
		d.CalculateStats()
		if d.Config.Verbose >= 0 {
			fmt.Printf("\r%6d %6d %13d %13d %13d %13d %8s %8.2f", d.Hstats.Loop, d.Hstats.Loops, d.Hstats.ReadBytes, d.Hstats.MaxBytes, d.Hstats.RemainingFiles, d.Hstats.HashedFiles, d.Hstats.Left, d.Hstats.MbPerSecond)
			fmt.Println("")
		}
		/*
			if !d.RegenHashmap() {
				fmt.Println("\nNo files left to hash")
				break
			}
		*/
	}
	d.Hstats.Duration = time.Since(d.Hstats.Start)
	if len(d.Hashes) < 1 {
		return false
	}
	return continueReading
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
	d.Hstats.Start = time.Now()

	if !d.LoadHashmap() {
		d.Hstats.Duration = time.Since(d.Hstats.Start)
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

// ProgressHeader @todo
func (d *Dupless) ProgressHeader() {
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

// DoReports @todo
func (d *Dupless) DoReports() {
	if d.Config.ErrorReport {
		d.ReportErrors()
	}
	if d.Config.IgnoredReport {
		d.ReportIgnored()
	}
	if d.Config.DirReport {
		d.ReportByDir()
	}
	if d.Config.SizeReport {
		d.ReportBySize()
	}
}

// Footer @todo
func (d *Dupless) Footer() {
	d.CalculateStats()
	fmt.Printf("\nFound %d matching files in %s\n", len(d.Files), d.Fstats.Duration)
	fmt.Printf("Found %d duplicate files in %s\n", d.Hstats.HashedFiles, d.Hstats.Duration)
	fmt.Printf("Total elapsed time: %s\n", d.Fstats.Duration+d.Hstats.Duration)
}

func usage() {
	fmt.Fprintf(os.Stderr, "\nUsage: %s [options] path [path2] ...\nOptions:\n\n", os.Args[0])
	flag.PrintDefaults()
}

func main() {
	basename := filepath.Base(os.Args[0])
	progname := strings.TrimSuffix(basename, filepath.Ext(basename))

	var d Dupless

	d.Init()

	fmt.Printf("%s: Version %s (%s)\n", progname, version.VERSION, version.GITCOMMIT)
	fmt.Printf("Built with %s for %s/%s (%d CPUs/%d GOMAXPROCS)\n",
		runtime.Version(),
		runtime.GOOS,
		runtime.GOARCH,
		runtime.NumCPU(),
		runtime.GOMAXPROCS(-1))

	if len(flag.Args()) == 0 || d.Config.Help {
		usage()
		return
	}

	d.ProcessArgs()

	d.ProgressHeader()

	if !d.FindFiles() {
		fmt.Println("No matching files found")
		return
	}

	ok := d.GetHashes()

	if ok {
		d.DoReports()
	} else {
		fmt.Println("No duplicate files found")
	}

	d.Footer()
}
