package report

import (
	"fmt"
	"sort"

	"github.com/rasa/godupless/dupless"
)

// IgnoredReport @todo
type IgnoredReport struct {
	d *dupless.Dupless
}

// NewIgnoredReport @todo
func NewIgnoredReport(d *dupless.Dupless) (r *IgnoredReport) {
	return &IgnoredReport{d: d}
}

// Run @todo
func (r *IgnoredReport) Run() {
	d := r.d
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
