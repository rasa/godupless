package report

import (
	"fmt"
	"github.com/rasa/godupless/dupless"
)

// DiffReport @todo
type DiffReport struct {
	d *dupless.Dupless
}

// NewDiffReport @todo
func NewDiffReport(d *dupless.Dupless) (r *DiffReport) {
	return &DiffReport{d: d}
}

// Run @todo
func (r *DiffReport) Run() {
	//d := r.d
	fmt.Print("\nDirectory Difference Report By Directory\n\n")
}
