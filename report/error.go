package report

import (
	"fmt"
	"sort"

	"github.com/rasa/godupless/dupless"
)

// ErrorReport @todo
type ErrorReport struct {
	d *dupless.Dupless
}

// NewErrorReport @todo
func NewErrorReport(d *dupless.Dupless) (r *ErrorReport) {
	return &ErrorReport{d: d}
}

// Run @todo
func (r *ErrorReport) Run() {
	d := r.d
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
		for _, errorFiles := range hashmap {
			fmt.Printf("%s: (%d bytes)\n", errorFiles.Err, size)
			for i, f := range errorFiles.Files {
				fmt.Printf("  %d: %s\n", i, f.Path())
			}
		}
	}
}
