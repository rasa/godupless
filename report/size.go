package report

import (
	"fmt"
	"sort"

	"github.com/rasa/godupless/dupless"
	"github.com/rasa/godupless/types"
)

// SizeReport @todo
type SizeReport struct {
	d *dupless.Dupless
}

// NewSizeReport @todo
func NewSizeReport(d *dupless.Dupless) (r *SizeReport) {
	return &SizeReport{d: d}
}

// Run @todo
func (r *SizeReport) Run() {
	d := r.d
	fmt.Print("\nDuplication Report By Size/Paths\n\n")

	i := 0
	sizes := make([]uint64, len(d.Dups))
	for size := range d.Dups {
		sizes[i] = size
		i++
	}
	sort.Sort(sort.Reverse(types.Uint64Slice(sizes)))

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
