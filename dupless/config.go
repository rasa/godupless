package dupless

import (
	"fmt"
	"os"
)

const (
	// MinChunk minimum chunk size, should be a power of 2
	MinChunk = 2 //2 << 12 // 4K
	// MaxChunk maximum chunk size, should be a power of 2
	MaxChunk = 2 << 29 // 1G
	// MinMinFiles Mininum number of the minimum files to compare to hash
	MinMinFiles = 2
)

// Config Config settings for Dupless class
type Config struct {
	//Cache string
	// Chunk @todo
	Chunk uint
	//Separator string
	// Exclude @todo
	Exclude string
	//Extra bool
	// Freq @todo
	Freq uint
	// Hash @todo
	Hash string
	// Help @todo
	Help bool
	// Iexclude @todo
	Iexclude string
	// Mask @todo
	Mask string
	// @todo add an imask option?
	// MinDirLength @todo
	MinDirLength uint
	// MinFiles @todo
	MinFiles uint
	// MinSize @todo
	MinSize uint64
	// Recursive @todo
	Recursive bool
	// Utc bool
	// Verbose @todo
	Verbose int
}

// DefaultConfig contains the default (initial) config settings
var DefaultConfig = Config{
	// Cache: "godupless.cache",
	Chunk: 2 << 16, // 2<<16 = 2^17 = 131,072
	// Separator: ",",
	//Exclude: "",
	// Extra: false,
	Freq: 100,
	Hash: DefaultHashFormat,
	//Help: false,
	//Iexclude: "",
	//Mask: "",
	MinFiles: 2,
	MinSize:  2 << 20, // 2<<20 = 2^21 = 2,097,152
	//Recursive: false,
	// Utc: false,
	//Verbose: 0,
}

// Validate @todo
func (c *Config) Validate() {
	if c.Chunk < MinChunk || c.Chunk > MaxChunk {
		fmt.Printf("Chunk must be between %d and %d", MinChunk, MaxChunk)
		os.Exit(1)
	}

	_ = GetHasher(c.Hash)

	if c.MinFiles < MinMinFiles {
		fmt.Printf("Mininum files to compare must be %d or greater", MinMinFiles)
		os.Exit(1)
	}
}
