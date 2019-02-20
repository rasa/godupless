package dupless

import (
	"hash"
	"os"
	"reflect"
	"testing"

	"github.com/rasa/godupless/file"
	"golang.org/x/text/message"
)

func TestDupless(t *testing.T) {
	t.Logf("TestDupless")
}

func Test_highway64New(t *testing.T) {
	tests := []struct {
		name string
		want hash.Hash
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := highway64New(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("highway64New() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_highway128New(t *testing.T) {
	tests := []struct {
		name string
		want hash.Hash
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := highway128New(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("highway128New() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_highway256New(t *testing.T) {
	tests := []struct {
		name string
		want hash.Hash
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := highway256New(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("highway256New() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_xxhashNew(t *testing.T) {
	tests := []struct {
		name string
		want hash.Hash
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := xxhashNew(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("xxhashNew() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDupless_Init(t *testing.T) {
	type fields struct {
		Config      Config
		Fstats      FileStats
		Hstats      HashStats
		P           *message.Printer
		Args        []string
		Excludes    []string
		Masks       []string
		Path        string
		Dev         string
		LastDev     string
		Volume      string
		HashFunc    func() hash.Hash
		ErrorDirs   map[string][]*ErrorRec
		IgnoredDirs map[string][]*IgnoredRec
		Files       map[string]*file.File
		Uniques     map[string][]string
		Sizes       map[uint64]map[string]*file.File
		Hashes      map[uint64]map[string][]*file.File
		Dups        map[uint64]map[string][]*file.File
		Errors      map[uint64]map[string]*ErrorFiles
	}
	tests := []struct {
		name   string
		fields fields
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			d := &Dupless{
				Config:      tt.fields.Config,
				Fstats:      tt.fields.Fstats,
				Hstats:      tt.fields.Hstats,
				P:           tt.fields.P,
				Args:        tt.fields.Args,
				Excludes:    tt.fields.Excludes,
				Masks:       tt.fields.Masks,
				Path:        tt.fields.Path,
				Dev:         tt.fields.Dev,
				LastDev:     tt.fields.LastDev,
				Volume:      tt.fields.Volume,
				HashFunc:    tt.fields.HashFunc,
				ErrorDirs:   tt.fields.ErrorDirs,
				IgnoredDirs: tt.fields.IgnoredDirs,
				Files:       tt.fields.Files,
				Uniques:     tt.fields.Uniques,
				Sizes:       tt.fields.Sizes,
				Hashes:      tt.fields.Hashes,
				Dups:        tt.fields.Dups,
				Errors:      tt.fields.Errors,
			}
			d.Init()
		})
	}
}

func TestDupless_AddError(t *testing.T) {
	type fields struct {
		Config      Config
		Fstats      FileStats
		Hstats      HashStats
		P           *message.Printer
		Args        []string
		Excludes    []string
		Masks       []string
		Path        string
		Dev         string
		LastDev     string
		Volume      string
		HashFunc    func() hash.Hash
		ErrorDirs   map[string][]*ErrorRec
		IgnoredDirs map[string][]*IgnoredRec
		Files       map[string]*file.File
		Uniques     map[string][]string
		Sizes       map[uint64]map[string]*file.File
		Hashes      map[uint64]map[string][]*file.File
		Dups        map[uint64]map[string][]*file.File
		Errors      map[uint64]map[string]*ErrorFiles
	}
	type args struct {
		path string
		s    string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			d := &Dupless{
				Config:      tt.fields.Config,
				Fstats:      tt.fields.Fstats,
				Hstats:      tt.fields.Hstats,
				P:           tt.fields.P,
				Args:        tt.fields.Args,
				Excludes:    tt.fields.Excludes,
				Masks:       tt.fields.Masks,
				Path:        tt.fields.Path,
				Dev:         tt.fields.Dev,
				LastDev:     tt.fields.LastDev,
				Volume:      tt.fields.Volume,
				HashFunc:    tt.fields.HashFunc,
				ErrorDirs:   tt.fields.ErrorDirs,
				IgnoredDirs: tt.fields.IgnoredDirs,
				Files:       tt.fields.Files,
				Uniques:     tt.fields.Uniques,
				Sizes:       tt.fields.Sizes,
				Hashes:      tt.fields.Hashes,
				Dups:        tt.fields.Dups,
				Errors:      tt.fields.Errors,
			}
			d.AddError(tt.args.path, tt.args.s)
		})
	}
}

func TestDupless_AddIgnore(t *testing.T) {
	type fields struct {
		Config      Config
		Fstats      FileStats
		Hstats      HashStats
		P           *message.Printer
		Args        []string
		Excludes    []string
		Masks       []string
		Path        string
		Dev         string
		LastDev     string
		Volume      string
		HashFunc    func() hash.Hash
		ErrorDirs   map[string][]*ErrorRec
		IgnoredDirs map[string][]*IgnoredRec
		Files       map[string]*file.File
		Uniques     map[string][]string
		Sizes       map[uint64]map[string]*file.File
		Hashes      map[uint64]map[string][]*file.File
		Dups        map[uint64]map[string][]*file.File
		Errors      map[uint64]map[string]*ErrorFiles
	}
	type args struct {
		path string
		typ  string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			d := &Dupless{
				Config:      tt.fields.Config,
				Fstats:      tt.fields.Fstats,
				Hstats:      tt.fields.Hstats,
				P:           tt.fields.P,
				Args:        tt.fields.Args,
				Excludes:    tt.fields.Excludes,
				Masks:       tt.fields.Masks,
				Path:        tt.fields.Path,
				Dev:         tt.fields.Dev,
				LastDev:     tt.fields.LastDev,
				Volume:      tt.fields.Volume,
				HashFunc:    tt.fields.HashFunc,
				ErrorDirs:   tt.fields.ErrorDirs,
				IgnoredDirs: tt.fields.IgnoredDirs,
				Files:       tt.fields.Files,
				Uniques:     tt.fields.Uniques,
				Sizes:       tt.fields.Sizes,
				Hashes:      tt.fields.Hashes,
				Dups:        tt.fields.Dups,
				Errors:      tt.fields.Errors,
			}
			d.AddIgnore(tt.args.path, tt.args.typ)
		})
	}
}

func TestDupless_AddDups(t *testing.T) {
	type fields struct {
		Config      Config
		Fstats      FileStats
		Hstats      HashStats
		P           *message.Printer
		Args        []string
		Excludes    []string
		Masks       []string
		Path        string
		Dev         string
		LastDev     string
		Volume      string
		HashFunc    func() hash.Hash
		ErrorDirs   map[string][]*ErrorRec
		IgnoredDirs map[string][]*IgnoredRec
		Files       map[string]*file.File
		Uniques     map[string][]string
		Sizes       map[uint64]map[string]*file.File
		Hashes      map[uint64]map[string][]*file.File
		Dups        map[uint64]map[string][]*file.File
		Errors      map[uint64]map[string]*ErrorFiles
	}
	type args struct {
		files []*file.File
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			d := &Dupless{
				Config:      tt.fields.Config,
				Fstats:      tt.fields.Fstats,
				Hstats:      tt.fields.Hstats,
				P:           tt.fields.P,
				Args:        tt.fields.Args,
				Excludes:    tt.fields.Excludes,
				Masks:       tt.fields.Masks,
				Path:        tt.fields.Path,
				Dev:         tt.fields.Dev,
				LastDev:     tt.fields.LastDev,
				Volume:      tt.fields.Volume,
				HashFunc:    tt.fields.HashFunc,
				ErrorDirs:   tt.fields.ErrorDirs,
				IgnoredDirs: tt.fields.IgnoredDirs,
				Files:       tt.fields.Files,
				Uniques:     tt.fields.Uniques,
				Sizes:       tt.fields.Sizes,
				Hashes:      tt.fields.Hashes,
				Dups:        tt.fields.Dups,
				Errors:      tt.fields.Errors,
			}
			d.AddDups(tt.args.files)
		})
	}
}

func TestDupless_ReadSize(t *testing.T) {
	type fields struct {
		Config      Config
		Fstats      FileStats
		Hstats      HashStats
		P           *message.Printer
		Args        []string
		Excludes    []string
		Masks       []string
		Path        string
		Dev         string
		LastDev     string
		Volume      string
		HashFunc    func() hash.Hash
		ErrorDirs   map[string][]*ErrorRec
		IgnoredDirs map[string][]*IgnoredRec
		Files       map[string]*file.File
		Uniques     map[string][]string
		Sizes       map[uint64]map[string]*file.File
		Hashes      map[uint64]map[string][]*file.File
		Dups        map[uint64]map[string][]*file.File
		Errors      map[uint64]map[string]*ErrorFiles
	}
	type args struct {
		size uint64
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			d := &Dupless{
				Config:      tt.fields.Config,
				Fstats:      tt.fields.Fstats,
				Hstats:      tt.fields.Hstats,
				P:           tt.fields.P,
				Args:        tt.fields.Args,
				Excludes:    tt.fields.Excludes,
				Masks:       tt.fields.Masks,
				Path:        tt.fields.Path,
				Dev:         tt.fields.Dev,
				LastDev:     tt.fields.LastDev,
				Volume:      tt.fields.Volume,
				HashFunc:    tt.fields.HashFunc,
				ErrorDirs:   tt.fields.ErrorDirs,
				IgnoredDirs: tt.fields.IgnoredDirs,
				Files:       tt.fields.Files,
				Uniques:     tt.fields.Uniques,
				Sizes:       tt.fields.Sizes,
				Hashes:      tt.fields.Hashes,
				Dups:        tt.fields.Dups,
				Errors:      tt.fields.Errors,
			}
			d.ReadSize(tt.args.size)
		})
	}
}

func TestDupless_CalculateStats(t *testing.T) {
	type fields struct {
		Config      Config
		Fstats      FileStats
		Hstats      HashStats
		P           *message.Printer
		Args        []string
		Excludes    []string
		Masks       []string
		Path        string
		Dev         string
		LastDev     string
		Volume      string
		HashFunc    func() hash.Hash
		ErrorDirs   map[string][]*ErrorRec
		IgnoredDirs map[string][]*IgnoredRec
		Files       map[string]*file.File
		Uniques     map[string][]string
		Sizes       map[uint64]map[string]*file.File
		Hashes      map[uint64]map[string][]*file.File
		Dups        map[uint64]map[string][]*file.File
		Errors      map[uint64]map[string]*ErrorFiles
	}
	tests := []struct {
		name   string
		fields fields
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			d := &Dupless{
				Config:      tt.fields.Config,
				Fstats:      tt.fields.Fstats,
				Hstats:      tt.fields.Hstats,
				P:           tt.fields.P,
				Args:        tt.fields.Args,
				Excludes:    tt.fields.Excludes,
				Masks:       tt.fields.Masks,
				Path:        tt.fields.Path,
				Dev:         tt.fields.Dev,
				LastDev:     tt.fields.LastDev,
				Volume:      tt.fields.Volume,
				HashFunc:    tt.fields.HashFunc,
				ErrorDirs:   tt.fields.ErrorDirs,
				IgnoredDirs: tt.fields.IgnoredDirs,
				Files:       tt.fields.Files,
				Uniques:     tt.fields.Uniques,
				Sizes:       tt.fields.Sizes,
				Hashes:      tt.fields.Hashes,
				Dups:        tt.fields.Dups,
				Errors:      tt.fields.Errors,
			}
			d.CalculateStats()
		})
	}
}

func TestDupless_ReadFiles(t *testing.T) {
	type fields struct {
		Config      Config
		Fstats      FileStats
		Hstats      HashStats
		P           *message.Printer
		Args        []string
		Excludes    []string
		Masks       []string
		Path        string
		Dev         string
		LastDev     string
		Volume      string
		HashFunc    func() hash.Hash
		ErrorDirs   map[string][]*ErrorRec
		IgnoredDirs map[string][]*IgnoredRec
		Files       map[string]*file.File
		Uniques     map[string][]string
		Sizes       map[uint64]map[string]*file.File
		Hashes      map[uint64]map[string][]*file.File
		Dups        map[uint64]map[string][]*file.File
		Errors      map[uint64]map[string]*ErrorFiles
	}
	tests := []struct {
		name   string
		fields fields
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			d := &Dupless{
				Config:      tt.fields.Config,
				Fstats:      tt.fields.Fstats,
				Hstats:      tt.fields.Hstats,
				P:           tt.fields.P,
				Args:        tt.fields.Args,
				Excludes:    tt.fields.Excludes,
				Masks:       tt.fields.Masks,
				Path:        tt.fields.Path,
				Dev:         tt.fields.Dev,
				LastDev:     tt.fields.LastDev,
				Volume:      tt.fields.Volume,
				HashFunc:    tt.fields.HashFunc,
				ErrorDirs:   tt.fields.ErrorDirs,
				IgnoredDirs: tt.fields.IgnoredDirs,
				Files:       tt.fields.Files,
				Uniques:     tt.fields.Uniques,
				Sizes:       tt.fields.Sizes,
				Hashes:      tt.fields.Hashes,
				Dups:        tt.fields.Dups,
				Errors:      tt.fields.Errors,
			}
			d.ReadFiles()
		})
	}
}

func TestDupless_LoadHashmap(t *testing.T) {
	type fields struct {
		Config      Config
		Fstats      FileStats
		Hstats      HashStats
		P           *message.Printer
		Args        []string
		Excludes    []string
		Masks       []string
		Path        string
		Dev         string
		LastDev     string
		Volume      string
		HashFunc    func() hash.Hash
		ErrorDirs   map[string][]*ErrorRec
		IgnoredDirs map[string][]*IgnoredRec
		Files       map[string]*file.File
		Uniques     map[string][]string
		Sizes       map[uint64]map[string]*file.File
		Hashes      map[uint64]map[string][]*file.File
		Dups        map[uint64]map[string][]*file.File
		Errors      map[uint64]map[string]*ErrorFiles
	}
	tests := []struct {
		name   string
		fields fields
		want   bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			d := &Dupless{
				Config:      tt.fields.Config,
				Fstats:      tt.fields.Fstats,
				Hstats:      tt.fields.Hstats,
				P:           tt.fields.P,
				Args:        tt.fields.Args,
				Excludes:    tt.fields.Excludes,
				Masks:       tt.fields.Masks,
				Path:        tt.fields.Path,
				Dev:         tt.fields.Dev,
				LastDev:     tt.fields.LastDev,
				Volume:      tt.fields.Volume,
				HashFunc:    tt.fields.HashFunc,
				ErrorDirs:   tt.fields.ErrorDirs,
				IgnoredDirs: tt.fields.IgnoredDirs,
				Files:       tt.fields.Files,
				Uniques:     tt.fields.Uniques,
				Sizes:       tt.fields.Sizes,
				Hashes:      tt.fields.Hashes,
				Dups:        tt.fields.Dups,
				Errors:      tt.fields.Errors,
			}
			if got := d.LoadHashmap(); got != tt.want {
				t.Errorf("Dupless.LoadHashmap() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDupless_GetHashes(t *testing.T) {
	type fields struct {
		Config      Config
		Fstats      FileStats
		Hstats      HashStats
		P           *message.Printer
		Args        []string
		Excludes    []string
		Masks       []string
		Path        string
		Dev         string
		LastDev     string
		Volume      string
		HashFunc    func() hash.Hash
		ErrorDirs   map[string][]*ErrorRec
		IgnoredDirs map[string][]*IgnoredRec
		Files       map[string]*file.File
		Uniques     map[string][]string
		Sizes       map[uint64]map[string]*file.File
		Hashes      map[uint64]map[string][]*file.File
		Dups        map[uint64]map[string][]*file.File
		Errors      map[uint64]map[string]*ErrorFiles
	}
	tests := []struct {
		name   string
		fields fields
		want   bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			d := &Dupless{
				Config:      tt.fields.Config,
				Fstats:      tt.fields.Fstats,
				Hstats:      tt.fields.Hstats,
				P:           tt.fields.P,
				Args:        tt.fields.Args,
				Excludes:    tt.fields.Excludes,
				Masks:       tt.fields.Masks,
				Path:        tt.fields.Path,
				Dev:         tt.fields.Dev,
				LastDev:     tt.fields.LastDev,
				Volume:      tt.fields.Volume,
				HashFunc:    tt.fields.HashFunc,
				ErrorDirs:   tt.fields.ErrorDirs,
				IgnoredDirs: tt.fields.IgnoredDirs,
				Files:       tt.fields.Files,
				Uniques:     tt.fields.Uniques,
				Sizes:       tt.fields.Sizes,
				Hashes:      tt.fields.Hashes,
				Dups:        tt.fields.Dups,
				Errors:      tt.fields.Errors,
			}
			if got := d.GetHashes(); got != tt.want {
				t.Errorf("Dupless.GetHashes() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDupless_Progress(t *testing.T) {
	type fields struct {
		Config      Config
		Fstats      FileStats
		Hstats      HashStats
		P           *message.Printer
		Args        []string
		Excludes    []string
		Masks       []string
		Path        string
		Dev         string
		LastDev     string
		Volume      string
		HashFunc    func() hash.Hash
		ErrorDirs   map[string][]*ErrorRec
		IgnoredDirs map[string][]*IgnoredRec
		Files       map[string]*file.File
		Uniques     map[string][]string
		Sizes       map[uint64]map[string]*file.File
		Hashes      map[uint64]map[string][]*file.File
		Dups        map[uint64]map[string][]*file.File
		Errors      map[uint64]map[string]*ErrorFiles
	}
	type args struct {
		force bool
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			d := &Dupless{
				Config:      tt.fields.Config,
				Fstats:      tt.fields.Fstats,
				Hstats:      tt.fields.Hstats,
				P:           tt.fields.P,
				Args:        tt.fields.Args,
				Excludes:    tt.fields.Excludes,
				Masks:       tt.fields.Masks,
				Path:        tt.fields.Path,
				Dev:         tt.fields.Dev,
				LastDev:     tt.fields.LastDev,
				Volume:      tt.fields.Volume,
				HashFunc:    tt.fields.HashFunc,
				ErrorDirs:   tt.fields.ErrorDirs,
				IgnoredDirs: tt.fields.IgnoredDirs,
				Files:       tt.fields.Files,
				Uniques:     tt.fields.Uniques,
				Sizes:       tt.fields.Sizes,
				Hashes:      tt.fields.Hashes,
				Dups:        tt.fields.Dups,
				Errors:      tt.fields.Errors,
			}
			d.Progress(tt.args.force)
		})
	}
}

func TestDupless_Visit(t *testing.T) {
	type fields struct {
		Config      Config
		Fstats      FileStats
		Hstats      HashStats
		P           *message.Printer
		Args        []string
		Excludes    []string
		Masks       []string
		Path        string
		Dev         string
		LastDev     string
		Volume      string
		HashFunc    func() hash.Hash
		ErrorDirs   map[string][]*ErrorRec
		IgnoredDirs map[string][]*IgnoredRec
		Files       map[string]*file.File
		Uniques     map[string][]string
		Sizes       map[uint64]map[string]*file.File
		Hashes      map[uint64]map[string][]*file.File
		Dups        map[uint64]map[string][]*file.File
		Errors      map[uint64]map[string]*ErrorFiles
	}
	type args struct {
		path string
		fi   os.FileInfo
		err  error
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			d := &Dupless{
				Config:      tt.fields.Config,
				Fstats:      tt.fields.Fstats,
				Hstats:      tt.fields.Hstats,
				P:           tt.fields.P,
				Args:        tt.fields.Args,
				Excludes:    tt.fields.Excludes,
				Masks:       tt.fields.Masks,
				Path:        tt.fields.Path,
				Dev:         tt.fields.Dev,
				LastDev:     tt.fields.LastDev,
				Volume:      tt.fields.Volume,
				HashFunc:    tt.fields.HashFunc,
				ErrorDirs:   tt.fields.ErrorDirs,
				IgnoredDirs: tt.fields.IgnoredDirs,
				Files:       tt.fields.Files,
				Uniques:     tt.fields.Uniques,
				Sizes:       tt.fields.Sizes,
				Hashes:      tt.fields.Hashes,
				Dups:        tt.fields.Dups,
				Errors:      tt.fields.Errors,
			}
			if err := d.Visit(tt.args.path, tt.args.fi, tt.args.err); (err != nil) != tt.wantErr {
				t.Errorf("Dupless.Visit() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestDupless_ResetCounters(t *testing.T) {
	type fields struct {
		Config      Config
		Fstats      FileStats
		Hstats      HashStats
		P           *message.Printer
		Args        []string
		Excludes    []string
		Masks       []string
		Path        string
		Dev         string
		LastDev     string
		Volume      string
		HashFunc    func() hash.Hash
		ErrorDirs   map[string][]*ErrorRec
		IgnoredDirs map[string][]*IgnoredRec
		Files       map[string]*file.File
		Uniques     map[string][]string
		Sizes       map[uint64]map[string]*file.File
		Hashes      map[uint64]map[string][]*file.File
		Dups        map[uint64]map[string][]*file.File
		Errors      map[uint64]map[string]*ErrorFiles
	}
	tests := []struct {
		name   string
		fields fields
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			d := &Dupless{
				Config:      tt.fields.Config,
				Fstats:      tt.fields.Fstats,
				Hstats:      tt.fields.Hstats,
				P:           tt.fields.P,
				Args:        tt.fields.Args,
				Excludes:    tt.fields.Excludes,
				Masks:       tt.fields.Masks,
				Path:        tt.fields.Path,
				Dev:         tt.fields.Dev,
				LastDev:     tt.fields.LastDev,
				Volume:      tt.fields.Volume,
				HashFunc:    tt.fields.HashFunc,
				ErrorDirs:   tt.fields.ErrorDirs,
				IgnoredDirs: tt.fields.IgnoredDirs,
				Files:       tt.fields.Files,
				Uniques:     tt.fields.Uniques,
				Sizes:       tt.fields.Sizes,
				Hashes:      tt.fields.Hashes,
				Dups:        tt.fields.Dups,
				Errors:      tt.fields.Errors,
			}
			d.ResetCounters()
		})
	}
}

func TestDupless_Header(t *testing.T) {
	type fields struct {
		Config      Config
		Fstats      FileStats
		Hstats      HashStats
		P           *message.Printer
		Args        []string
		Excludes    []string
		Masks       []string
		Path        string
		Dev         string
		LastDev     string
		Volume      string
		HashFunc    func() hash.Hash
		ErrorDirs   map[string][]*ErrorRec
		IgnoredDirs map[string][]*IgnoredRec
		Files       map[string]*file.File
		Uniques     map[string][]string
		Sizes       map[uint64]map[string]*file.File
		Hashes      map[uint64]map[string][]*file.File
		Dups        map[uint64]map[string][]*file.File
		Errors      map[uint64]map[string]*ErrorFiles
	}
	tests := []struct {
		name   string
		fields fields
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			d := &Dupless{
				Config:      tt.fields.Config,
				Fstats:      tt.fields.Fstats,
				Hstats:      tt.fields.Hstats,
				P:           tt.fields.P,
				Args:        tt.fields.Args,
				Excludes:    tt.fields.Excludes,
				Masks:       tt.fields.Masks,
				Path:        tt.fields.Path,
				Dev:         tt.fields.Dev,
				LastDev:     tt.fields.LastDev,
				Volume:      tt.fields.Volume,
				HashFunc:    tt.fields.HashFunc,
				ErrorDirs:   tt.fields.ErrorDirs,
				IgnoredDirs: tt.fields.IgnoredDirs,
				Files:       tt.fields.Files,
				Uniques:     tt.fields.Uniques,
				Sizes:       tt.fields.Sizes,
				Hashes:      tt.fields.Hashes,
				Dups:        tt.fields.Dups,
				Errors:      tt.fields.Errors,
			}
			d.Header()
		})
	}
}

func TestDupless_ProcessArgs(t *testing.T) {
	type fields struct {
		Config      Config
		Fstats      FileStats
		Hstats      HashStats
		P           *message.Printer
		Args        []string
		Excludes    []string
		Masks       []string
		Path        string
		Dev         string
		LastDev     string
		Volume      string
		HashFunc    func() hash.Hash
		ErrorDirs   map[string][]*ErrorRec
		IgnoredDirs map[string][]*IgnoredRec
		Files       map[string]*file.File
		Uniques     map[string][]string
		Sizes       map[uint64]map[string]*file.File
		Hashes      map[uint64]map[string][]*file.File
		Dups        map[uint64]map[string][]*file.File
		Errors      map[uint64]map[string]*ErrorFiles
	}
	tests := []struct {
		name   string
		fields fields
		want   bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			d := &Dupless{
				Config:      tt.fields.Config,
				Fstats:      tt.fields.Fstats,
				Hstats:      tt.fields.Hstats,
				P:           tt.fields.P,
				Args:        tt.fields.Args,
				Excludes:    tt.fields.Excludes,
				Masks:       tt.fields.Masks,
				Path:        tt.fields.Path,
				Dev:         tt.fields.Dev,
				LastDev:     tt.fields.LastDev,
				Volume:      tt.fields.Volume,
				HashFunc:    tt.fields.HashFunc,
				ErrorDirs:   tt.fields.ErrorDirs,
				IgnoredDirs: tt.fields.IgnoredDirs,
				Files:       tt.fields.Files,
				Uniques:     tt.fields.Uniques,
				Sizes:       tt.fields.Sizes,
				Hashes:      tt.fields.Hashes,
				Dups:        tt.fields.Dups,
				Errors:      tt.fields.Errors,
			}
			if got := d.ProcessArgs(); got != tt.want {
				t.Errorf("Dupless.ProcessArgs() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDupless_FindFiles(t *testing.T) {
	type fields struct {
		Config      Config
		Fstats      FileStats
		Hstats      HashStats
		P           *message.Printer
		Args        []string
		Excludes    []string
		Masks       []string
		Path        string
		Dev         string
		LastDev     string
		Volume      string
		HashFunc    func() hash.Hash
		ErrorDirs   map[string][]*ErrorRec
		IgnoredDirs map[string][]*IgnoredRec
		Files       map[string]*file.File
		Uniques     map[string][]string
		Sizes       map[uint64]map[string]*file.File
		Hashes      map[uint64]map[string][]*file.File
		Dups        map[uint64]map[string][]*file.File
		Errors      map[uint64]map[string]*ErrorFiles
	}
	tests := []struct {
		name   string
		fields fields
		want   bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			d := &Dupless{
				Config:      tt.fields.Config,
				Fstats:      tt.fields.Fstats,
				Hstats:      tt.fields.Hstats,
				P:           tt.fields.P,
				Args:        tt.fields.Args,
				Excludes:    tt.fields.Excludes,
				Masks:       tt.fields.Masks,
				Path:        tt.fields.Path,
				Dev:         tt.fields.Dev,
				LastDev:     tt.fields.LastDev,
				Volume:      tt.fields.Volume,
				HashFunc:    tt.fields.HashFunc,
				ErrorDirs:   tt.fields.ErrorDirs,
				IgnoredDirs: tt.fields.IgnoredDirs,
				Files:       tt.fields.Files,
				Uniques:     tt.fields.Uniques,
				Sizes:       tt.fields.Sizes,
				Hashes:      tt.fields.Hashes,
				Dups:        tt.fields.Dups,
				Errors:      tt.fields.Errors,
			}
			if got := d.FindFiles(); got != tt.want {
				t.Errorf("Dupless.FindFiles() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDupless_Footer(t *testing.T) {
	type fields struct {
		Config      Config
		Fstats      FileStats
		Hstats      HashStats
		P           *message.Printer
		Args        []string
		Excludes    []string
		Masks       []string
		Path        string
		Dev         string
		LastDev     string
		Volume      string
		HashFunc    func() hash.Hash
		ErrorDirs   map[string][]*ErrorRec
		IgnoredDirs map[string][]*IgnoredRec
		Files       map[string]*file.File
		Uniques     map[string][]string
		Sizes       map[uint64]map[string]*file.File
		Hashes      map[uint64]map[string][]*file.File
		Dups        map[uint64]map[string][]*file.File
		Errors      map[uint64]map[string]*ErrorFiles
	}
	tests := []struct {
		name   string
		fields fields
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			d := &Dupless{
				Config:      tt.fields.Config,
				Fstats:      tt.fields.Fstats,
				Hstats:      tt.fields.Hstats,
				P:           tt.fields.P,
				Args:        tt.fields.Args,
				Excludes:    tt.fields.Excludes,
				Masks:       tt.fields.Masks,
				Path:        tt.fields.Path,
				Dev:         tt.fields.Dev,
				LastDev:     tt.fields.LastDev,
				Volume:      tt.fields.Volume,
				HashFunc:    tt.fields.HashFunc,
				ErrorDirs:   tt.fields.ErrorDirs,
				IgnoredDirs: tt.fields.IgnoredDirs,
				Files:       tt.fields.Files,
				Uniques:     tt.fields.Uniques,
				Sizes:       tt.fields.Sizes,
				Hashes:      tt.fields.Hashes,
				Dups:        tt.fields.Dups,
				Errors:      tt.fields.Errors,
			}
			d.Footer()
		})
	}
}

func TestDupless_Run(t *testing.T) {
	type fields struct {
		Config      Config
		Fstats      FileStats
		Hstats      HashStats
		P           *message.Printer
		Args        []string
		Excludes    []string
		Masks       []string
		Path        string
		Dev         string
		LastDev     string
		Volume      string
		HashFunc    func() hash.Hash
		ErrorDirs   map[string][]*ErrorRec
		IgnoredDirs map[string][]*IgnoredRec
		Files       map[string]*file.File
		Uniques     map[string][]string
		Sizes       map[uint64]map[string]*file.File
		Hashes      map[uint64]map[string][]*file.File
		Dups        map[uint64]map[string][]*file.File
		Errors      map[uint64]map[string]*ErrorFiles
	}
	tests := []struct {
		name   string
		fields fields
		want   bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			d := &Dupless{
				Config:      tt.fields.Config,
				Fstats:      tt.fields.Fstats,
				Hstats:      tt.fields.Hstats,
				P:           tt.fields.P,
				Args:        tt.fields.Args,
				Excludes:    tt.fields.Excludes,
				Masks:       tt.fields.Masks,
				Path:        tt.fields.Path,
				Dev:         tt.fields.Dev,
				LastDev:     tt.fields.LastDev,
				Volume:      tt.fields.Volume,
				HashFunc:    tt.fields.HashFunc,
				ErrorDirs:   tt.fields.ErrorDirs,
				IgnoredDirs: tt.fields.IgnoredDirs,
				Files:       tt.fields.Files,
				Uniques:     tt.fields.Uniques,
				Sizes:       tt.fields.Sizes,
				Hashes:      tt.fields.Hashes,
				Dups:        tt.fields.Dups,
				Errors:      tt.fields.Errors,
			}
			if got := d.Run(); got != tt.want {
				t.Errorf("Dupless.Run() = %v, want %v", got, tt.want)
			}
		})
	}
}
