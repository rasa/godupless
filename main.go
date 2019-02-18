// Program godupless creates a report of duplicate files across multiple volumes
package main

import (
	"flag"
	"fmt"
	"github.com/rasa/godupless/dupless"
	"github.com/rasa/godupless/report"
	"github.com/rasa/godupless/version"
	"os"
	"path/filepath"
	"runtime"
	"strings"
)

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

	var d dupless.Dupless

	rv := d.Run()
	if !rv {
		if d.Config.Help {
			usage()
		}
		os.Exit(1)
	}

	if d.Config.ErrorReport {
		r := report.NewErrorReport(&d)
		r.Run()
	}
	if d.Config.IgnoredReport {
		r := report.NewIgnoredReport(&d)
		r.Run()
	}
	if d.Config.HardLinkReport {
		r := report.NewHardLinkReport(&d)
		r.Run()
	}
	if d.Config.DiffReport {
		r := report.NewDiffReport(&d)
		r.Run()
	}
	if d.Config.DirReport {
		r := report.NewDirReport(&d)
		r.Run()
	}
	if d.Config.SizeReport {
		r := report.NewSizeReport(&d)
		r.Run()
	}

	d.Footer()

	os.Exit(0)
}
