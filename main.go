// Program godupless prints the version and exits
package main

import (
	"crypto/md5"
	"crypto/sha1"
	"crypto/sha256"
	"crypto/sha512"
	"encoding/json"
	"flag"
	"fmt"
	"hash"
	"io"
	"log"
	"os"
	"path/filepath"
	"runtime"
	"sort"
	"strings"

	"github.com/rasa/godupless/version"
	"golang.org/x/text/language"
	"golang.org/x/text/message"
)

/*
2^1 	2
2^2 	4
2^3 	8
2^4 	16
2^5 	32
2^6 	64
2^7 	128
2^8 	256
2^9 	512
2^10	1_024
2^11	2_048
2^12	4_096
2^13	8_192
2^14	16_384
2^15	32_768
2^16	65_536
2^17	131_072
2^18	262_144
2^19	524_288
2^20	1_048_576
2^21	2_097_152
2^22	4_194_304
2^23	8_388_608
2^24	16_777_216
2^25	33_554_432
2^26	67_108_864
2^27	134_217_728
2^28	268_435_456
2^29	536_870_912
2^30	1_073_741_824
2^31	2_147_483_648
2^32	4_294_967_296
*/

const (
	// DefaultDirReport @todo
	DefaultDirReport = false
	// DefaultFrequency @todo
	DefaultFrequency = 100
	// DefaultHash @todo
	DefaultHash = "md5"
	// DefaultMask @todo
	DefaultMask = ""
	// DefaultMinFiles @todo
	DefaultMinFiles = 2
	// DefaultMinDirLength @todo
	DefaultMinDirLength = uint(2)
	// DefaultMinSize @todo
	DefaultMinSize = 2 << 24 // 16M
	// DefaultRecursive @todo
	DefaultRecursive = false
	// DefaultSizeReport @todo
	DefaultSizeReport = true
	// DefaultVerbose @todo
	DefaultVerbose = 0
)

// Dir @todo
type Dir struct {
	// Count @todo
	Count uint
	// Size @todo
	Size uint64
}

// DirCount @todo
type DirCount struct {
	// Dir @todo
	Dir string
	// Count @todo
	Count uint
}

// ErrorRec @todo
type ErrorRec struct {
	// Path @todo
	Path string
	// Error @todo
	Error string
}

// IgnoredRec @todo
type IgnoredRec struct {
	// Path @todo
	Path string
	// Type @todo
	Type string
}

// Dupless @todo
type Dupless struct {
	dirReport    bool
	freq         uint
	hash         string
	help         bool
	mask         string
	minDirLength uint
	minFiles     uint
	minSize      uint64
	sizeReport   bool
	recursive    bool
	verbose      uint

	lastVolume string
	volume     string
	// path        string
	p *message.Printer

	// volume stats:
	hits        uint
	skipped     uint
	directories uint
	matched     uint
	errors      uint
	ignored     uint
	dirs        map[string]*Dir
	errorDirs   map[string][]*ErrorRec
	ignoredDirs map[string][]*IgnoredRec
	dups        map[uint64]map[string][]string
}

// Uint64Slice @todo
type Uint64Slice []uint64

func (p Uint64Slice) Len() int           { return len(p) }
func (p Uint64Slice) Less(i, j int) bool { return p[i] < p[j] }
func (p Uint64Slice) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }

// Sort is a convenience method.
func (p Uint64Slice) Sort() { sort.Sort(p) }

func dump(s string, x interface{}) {
	if s != "" {
		fmt.Print(s)
	}
	if x == nil {
		return
	}

	b, err := json.MarshalIndent(x, "", "  ")
	if err != nil {
		log.Fatal("\nJSON marshaling error: ", err)
	}
	fmt.Print(string(b))
}

func getDir(path string, minDirLength uint) string {
	dir, _ := filepath.Split(path)
	if uint(len(dir)) <= minDirLength {
		return dir
	}
	return trimSuffix(dir, string(os.PathSeparator))
}

func getFileType(fi os.FileInfo) string {
	switch mode := fi.Mode(); {
	case mode.IsRegular():
		return "file"
	case mode.IsDir():
		return "directory"
	case mode&os.ModeSymlink != 0:
		return "symlink"
	case mode&os.ModeDevice != 0:
		return "device"
	case mode&os.ModeNamedPipe != 0:
		return "named pipe"
	case mode&os.ModeSocket != 0:
		return "socket"
	case mode&os.ModeIrregular != 0:
		return "irregular"
	}
	return "unknown type"
}

func substring(s string, start int, end int) string {
	startStrIdx := 0
	i := 0
	for j := range s {
		if i == start {
			startStrIdx = j
		}
		if i == end {
			return s[startStrIdx:j]
		}
		i++
	}
	return s[startStrIdx:]
}

func trimSuffix(s, suffix string) string {
	if strings.HasSuffix(s, suffix) {
		s = s[:len(s)-len(suffix)]
	}
	return s
}

func (d *Dupless) init() {
	minDirLength := DefaultMinDirLength
	if runtime.GOOS == "windows" {
		// c:/
		minDirLength = 3
	}

	flag.BoolVar(&d.dirReport, "dir_report", DefaultDirReport, "Report by directory")
	flag.UintVar(&d.freq, "frequency", DefaultFrequency, "Reporting frequency")
	flag.StringVar(&d.hash, "hash", DefaultHash, "Hash type")
	flag.BoolVar(&d.help, "help", false, "Display help")
	flag.StringVar(&d.mask, "mask", DefaultMask, "File mask")
	flag.UintVar(&d.minFiles, "min_files", DefaultMinFiles, "Minimum files")
	flag.UintVar(&d.minDirLength, "min_dir_len", minDirLength, "Minimum directory length")
	flag.Uint64Var(&d.minSize, "min_size", DefaultMinSize, "Minimum file size")
	flag.BoolVar(&d.sizeReport, "size_report", DefaultSizeReport, "Report by size")
	flag.BoolVar(&d.recursive, "recursive", DefaultRecursive, "Report directories recursively")
	flag.UintVar(&d.verbose, "verbose", DefaultVerbose, "Be more verbose")

	d.dirs = make(map[string]*Dir)
	d.dups = make(map[uint64]map[string][]string)
	d.errorDirs = make(map[string][]*ErrorRec)
	d.ignoredDirs = make(map[string][]*IgnoredRec)
	d.p = message.NewPrinter(language.English)
}

func (d *Dupless) addPath(path string, size uint64) {
	for {
		dir := getDir(path, d.minDirLength)
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
	if d.volume != d.lastVolume {
		if d.lastVolume != "" {
			d.skipped = 0
			d.directories = 0
			d.matched = 0
			d.errors = 0
			d.ignored = 0
			fmt.Println("")
		}
		d.lastVolume = d.volume
	}

	d.p.Printf("\r%11d %11d %11d %11d %11d %s", d.skipped, d.matched, d.directories, d.ignored, d.errors, d.volume)

	if final {
		fmt.Println("")
	}
}

func (d *Dupless) addError(path string, s string) {
	dir := getDir(path, d.minDirLength)
	errorRec := ErrorRec{Path: path, Error: s}
	d.errorDirs[dir] = append(d.errorDirs[dir], &errorRec)
	d.errors++
	if d.verbose > 0 {
		fmt.Fprintln(os.Stderr, s)
	}
}

func (d *Dupless) addIgnore(path string, typ string) {
	dir := getDir(path, d.minDirLength)
	IgnoredRec := IgnoredRec{Path: path, Type: typ}
	d.ignoredDirs[dir] = append(d.ignoredDirs[dir], &IgnoredRec)
	d.ignored++
	if d.verbose > 0 {
		s := fmt.Sprintf("Skipping '%s' as it is a %s", path, typ)
		fmt.Fprintf(os.Stderr, "\n%s\n", s)
	}
}

func (d *Dupless) reportByDir() {
	fmt.Printf("Duplication Report By Size/Directory\n\n")

	for size := range d.dups {
		for _, paths := range d.dups[size] {
			if len(paths) < 2 {
				continue
			}
			for _, path := range paths {
				d.addPath(path, size)
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

	totalSize := uint64(0)
	files := uint64(0)

	for _, size := range sizes {
		emitSize := true
		totalSize += size
		for i, dircount := range sizemap[size] {
			if emitSize {
				fmt.Printf("size: %d\n", size)
				emitSize = false
			}
			fmt.Printf("  %d: %v (%d)\n", i+1, dircount.Dir, dircount.Count)
			files += uint64(dircount.Count)
		}
	}
	d.p.Printf("\n%d files totaling %d bytes (%d bytes per file average)\n", files, totalSize, totalSize/files)
}

func (d *Dupless) reportBySize() {
	fmt.Printf("Duplication Report By Size/Paths\n\n")

	i := 0
	sizes := make([]uint64, len(d.dups))
	for size := range d.dups {
		sizes[i] = size
		i++
	}
	sort.Sort(sort.Reverse(Uint64Slice(sizes)))

	totalSize := uint64(0)
	files := uint64(0)

	for _, size := range sizes {
		for hash, paths := range d.dups[size] {
			if uint(len(paths)) < d.minFiles {
				continue
			}
			totalSize += size
			files += uint64(len(paths))
			d.p.Printf("Size: %d (%s)\n", size, hash)
			for i, path := range paths {
				fmt.Printf("  %d: %s\n", i+1, path)
			}
		}
	}

	d.p.Printf("\n%d files totaling %d bytes (%d bytes per file average)\n", files, totalSize, totalSize/files)
}

func (d *Dupless) reportIgnored() {
	if len(d.ignoredDirs) == 0 {
		return
	}

	fmt.Printf("Ignored Files/Directories Report\n\n")

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

func (d *Dupless) visit(path string, f os.FileInfo, err error) error {
	for {
		d.hits++
		if runtime.GOOS == "windows" {
			d.volume = substring(path, 0, 2)
		} else {
			/*
			   	// https://groups.google.com/forum/#!topic/golang-nuts/mu8XMmRXMOk
			   	// https://stackoverflow.com/q/19513874/1432614 :
			      mntpoint, err := os.Stat(mountpoint)
			       if err != nil {
			           if os.IsNotExist(err) {
			                   return false, nil
			           }
			           return false, err
			       }
			       parent, err := os.Stat(filepath.Join(mountpoint, ".."))
			       if err != nil {
			           return false, err
			       }
			       mntpointSt := mntpoint.Sys().(*syscall.Stat_t)
			       parentSt := parent.Sys().(*syscall.Stat_t)
			       return mntpointSt.Dev != parentSt.Dev, nil
			*/
		}
		fi, err := os.Lstat(path)
		if err != nil {
			s := fmt.Sprintf("Cannot lstat '%s': %s", path, err)
			d.addError(path, s)
			break
		}
		typ := getFileType(fi)

		if typ == "directory" {
			d.directories++
			break
		}
		if typ == "symlink" {
			d.ignored++
			break
		}
		if typ != "file" {
			d.addIgnore(path, typ)
			break
		}

		size := uint64(fi.Size())
		if size <= d.minSize {
			d.skipped++
			break
		}

		if d.mask != "" {
			_, file := filepath.Split(path)
			ok, err := filepath.Match(d.mask, file)
			if err != nil {
				d.errors++
				if d.verbose > 0 {
					fmt.Fprintf(os.Stderr, "\nCannot match '%s' using %s: %s\n", path, d.mask, err)
				}
				break
			}
			if !ok {
				d.skipped++
				break
			}
		}

		fh, err := os.Open(path)
		if err != nil {
			d.errors++
			if d.verbose > 0 {
				fmt.Fprintf(os.Stderr, "\nCannot open '%s': %s\n", path, err)
			}
			break
		}
		defer fh.Close()

		var h hash.Hash

		switch d.hash {
		case "md5", "MD5":
			h = md5.New()
		case "sha1", "SHA1":
			h = sha1.New()
		case "sha256", "SHA256":
			h = sha256.New()
		case "sha512", "SHA512":
			h = sha512.New()
		default:
			fmt.Fprintf(os.Stderr, "Unknown hash format: '%s'\n", d.hash)
			os.Exit(1)
		}

		_, err = io.Copy(h, fh)
		if err != nil {
			d.errors++
			if d.verbose > 0 {
				fmt.Fprintf(os.Stderr, "\nCannot read '%s': %s\n", path, err)
			}
			break
		}

		d.matched++

		hash := fmt.Sprintf("%x", h.Sum(nil))
		_, ok := d.dups[size]
		if !ok {
			d.dups[size] = make(map[string][]string)
		}
		_, ok = d.dups[size][hash]
		if !ok {
			d.dups[size][hash] = make([]string, 0)
		}
		d.dups[size][hash] = append(d.dups[size][hash], path)
		break
	}

	d.progress(false)
	return nil
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

	fmt.Printf("    skipped     matched directories     ignored      errors volume\n")
	fmt.Printf("----------- ----------- ----------- ----------- ----------- ------\n")

	for _, arg := range flag.Args() {
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
	if dupless.dirReport {
		dupless.reportByDir()
	}
	if dupless.sizeReport {
		dupless.reportBySize()
	}
}
