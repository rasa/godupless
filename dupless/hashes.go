package dupless

import (
	"crypto/md5"
	"crypto/sha1"
	"crypto/sha256"
	"crypto/sha512"
	"encoding/hex"
	"fmt"
	"hash"
	"os"
	"strings"

	"github.com/cespare/xxhash"
	"github.com/minio/highwayhash"
)

const (
	// @todo move to Config?
	// from https://github.com/minio/highwayhash/blob/master/examples_test.go#L17
	//                              1                   2                   3
	//             1 2 3 4 5 6 7 8 9 0 1 2 3 4 5 6 7 8 9 0 1 2 3 4 5 6 7 8 9 0 1 2
	hashString = "000102030405060708090A0B0C0D0E0FF0E0D0C0B0A090807060504030201000"
	// DefaultHashFormat @todo
	DefaultHashFormat = "highway"
)

var hashKey []byte

func highway64New() hash.Hash {
	h, err := highwayhash.New64(hashKey)
	if err != nil {
		panic(err)
	}
	return h
}

func highway128New() hash.Hash {
	h, err := highwayhash.New128(hashKey)
	if err != nil {
		panic(err)
	}
	return h
}

func highway256New() hash.Hash {
	h, err := highwayhash.New(hashKey)
	if err != nil {
		panic(err)
	}
	return h
}

func xxhashNew() hash.Hash {
	return xxhash.New()
}

// Hashmap list of hash methods
var Hashmap = map[string]func() hash.Hash{
	"highway":    highway256New,
	"highway64":  highway64New,
	"highway128": highway128New,
	"highway256": highway256New,
	"md5":        md5.New,
	"sha1":       sha1.New,
	"sha256":     sha256.New,
	"sha512":     sha512.New,
	"xxhash":     xxhashNew,
}

// Hashes @todo
func Hashes() string {
	var hashes []string
	for k := range Hashmap {
		hashes = append(hashes, k)
	}
	return strings.Join(hashes, ",")
}

func init() {
	var err error
	hashKey, err = hex.DecodeString(hashString)
	if err != nil {
		fmt.Printf("Invalid hash key %s: %s\n", hashString, err)
		os.Exit(1)
	}
}

// GetHasher @todo
func GetHasher(hash string) func() hash.Hash {
	hash = strings.ToLower(hash)
	hasher, ok := Hashmap[hash]
	if !ok {
		fmt.Println("Hash %s is not one of: ", hash, Hashes())
		os.Exit(1)
	}
	return hasher
}
