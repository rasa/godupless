package report

import (
	"fmt"
	"sort"

	"github.com/rasa/godupless/dupless"
	"github.com/rasa/godupless/file"
	"github.com/rasa/godupless/types"
	"github.com/rasa/godupless/util"
)

// DirReport @todo
type DirReport struct {
	d *dupless.Dupless
	// Dirs Dirs[dir] = *Dir{Count, Size}
	Dirs map[string]*Dir
}

// Dir Directory report record
type Dir struct {
	// Count number of duplicate files in the directory
	Count uint64
	// Size total size of the duplicate files in the directory
	Size uint64
}

// DirCount Used internally by Directory report function
type DirCount struct {
	// Dir contains the full of the directory containing one or more files that are duplicated
	Dir string
	// Count contains the number of files in directory that are duplicated in other directories
	Count uint64
}

// NewDirReport @todo
func NewDirReport(d *dupless.Dupless) (r *DirReport) {
	return &DirReport{d: d, Dirs: make(map[string]*Dir)}
}

// Run @todo
func (r *DirReport) Run() {
	d := r.d
	fmt.Print("\nDuplication Summary Report By Size/Directory\n\n")

	for size, hashmap := range d.Dups {
		for _, files := range hashmap {
			for _, f := range files {
				r.addPath(f.Path(), size)
			}
		}
	}

	sizemap := make(map[uint64][]*DirCount)

	for dir, dr := range r.Dirs {
		size := dr.Size
		/*_, ok := sizemap[size]
		if !ok {
			sizemap[size] = make([]*DirCount, 0)
		}*/
		sizemap[size] = append(sizemap[size], &DirCount{Dir: dir, Count: dr.Count})
	}

	i := 0
	sizes := make([]uint64, len(sizemap))
	for size := range sizemap {
		sizes[i] = size
		i++
	}
	sort.Sort(sort.Reverse(types.Uint64Slice(sizes)))

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

// addPath @todo
func (r *DirReport) addPath(path string, size uint64) {
	d := r.d
	for {
		dir := util.Dirname(path)
		if uint(len(dir)) < file.MinDirLength {
			return
		}
		if dir == path {
			return
		}
		_, ok := r.Dirs[dir]
		if !ok {
			r.Dirs[dir] = &Dir{Count: 1, Size: size}
		} else {
			r.Dirs[dir].Count++
			r.Dirs[dir].Size += size
		}
		if !d.Config.Recursive {
			return
		}
		path = dir
	}
}
