package report

import (
	"fmt"

	"github.com/rasa/godupless/dupless"
	// "github.com/rasa/godupless/file"
	//"k8s.io/apimachinery/pkg/util/sets"
)

// HardLinkReport @todo
type HardLinkReport struct {
	d *dupless.Dupless
}

// NewHardLinkReport @todo
func NewHardLinkReport(d *dupless.Dupless) (r *HardLinkReport) {
	return &HardLinkReport{d: d}
}

/*
@todo
unique report:
	scan d.Dups:
		scan D.Uniques:
			report if name/mtime/ctime/mode/uid/gid is different
*/

// Run @todo
func (r *HardLinkReport) Run() {
	// d := r.d
	fmt.Print("\nHard Link Report\n\n")
}
