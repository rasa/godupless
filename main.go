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

type config struct {
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

var defaultConfig = config{
	//DiffReport: false,
	//DirReport: false,
	//ErrorReport: false,
	//HardLinkReport: false,
	//IgnoredReport: false,
	SizeReport: true,
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

	c := dupless.DefaultConfig
	d := dupless.New(c)

	chunk := fmt.Sprintf("Hash chunk size (%d to %d)", dupless.MinChunk, dupless.MaxChunk)

	hash := "Hash type: " + dupless.Hashes()

	// flag.StringVar(&d.Config.Cache, "cache", d.Config.Cache, "Cache filename")
	flag.UintVar(&c.Chunk, "chunk", c.Chunk, chunk)
	flag.StringVar(&c.Exclude, "exclude", c.Exclude, "Regex(s) of directories/files to exclude, separated by |")
	flag.StringVar(&c.Iexclude, "iexclude", c.Iexclude, "Regex(s) of directories/files to exclude, separated by |")
	//flag.BoolVar(&c.Extra, "extra", c.Extra, "Cache extra attributes")
	flag.UintVar(&c.Freq, "frequency", c.Freq, "Reporting frequency")
	flag.StringVar(&c.Hash, "hash", c.Hash, hash)
	flag.BoolVar(&c.Help, "help", c.Help, "Display help")
	flag.StringVar(&c.Mask, "mask", c.Mask, "File mask(s), seperated by |")
	flag.UintVar(&c.MinFiles, "min_files", c.MinFiles, "Minimum files to compare")
	flag.Uint64Var(&c.MinSize, "min_size", c.MinSize, "Minimum file size")
	flag.BoolVar(&c.Recursive, "recursive", c.Recursive, "Report directories recursively")
	//flag.StringVar(&c.Seperator, "separator", c.Seperator, "Field separator")
	// flag.BoolVar(&c.Utc, "utc", c.Utc, "Report times in UTC")
	flag.IntVar(&c.Verbose, "verbose", c.Verbose, "Increase log verbosity")

	flag.BoolVar(&defaultConfig.DiffReport, "diff_report", defaultConfig.DiffReport, "Report on differences between directories containing duplicate files")
	flag.BoolVar(&defaultConfig.DirReport, "dir_report", defaultConfig.DirReport, "Report summary of duplicates by directory")
	flag.BoolVar(&defaultConfig.ErrorReport, "error_report", defaultConfig.ErrorReport, "Report of errors")
	flag.BoolVar(&defaultConfig.HardLinkReport, "hard_link_report", defaultConfig.HardLinkReport, "Report on hard link differences")
	flag.BoolVar(&defaultConfig.IgnoredReport, "ignored_report", defaultConfig.IgnoredReport, "Report of ignored files")
	flag.BoolVar(&defaultConfig.SizeReport, "size_report", defaultConfig.SizeReport, "Report duplicates by size")

	flag.Parse()

	c.Validate()

	rv := d.Run()
	if !rv {
		if d.Config.Help {
			usage()
		}
		os.Exit(1)
	}

	if defaultConfig.ErrorReport {
		r := report.NewErrorReport(d)
		r.Run()
	}
	if defaultConfig.IgnoredReport {
		r := report.NewIgnoredReport(d)
		r.Run()
	}
	if defaultConfig.HardLinkReport {
		r := report.NewHardLinkReport(d)
		r.Run()
	}
	if defaultConfig.DiffReport {
		r := report.NewDiffReport(d)
		r.Run()
	}
	if defaultConfig.DirReport {
		r := report.NewDirReport(d)
		r.Run()
	}
	if defaultConfig.SizeReport {
		r := report.NewSizeReport(d)
		r.Run()
	}

	d.Footer()

	os.Exit(0)
}
