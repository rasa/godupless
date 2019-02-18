package util

import (
	"bufio"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"path/filepath"
	//"runtime"
	"strings"
	"syscall"
	"time"
)

// Basename @todo rename to Base
func Basename(path string) string {
	return filepath.Base(path)
}

// Dirname @todo rename to Dir
func Dirname(path string) string {
	return filepath.Dir(path)
}

// Dump @todo
func Dump(s string, x interface{}) {
	if s != "" {
		fmt.Print(s)
	}
	if x == nil {
		return
	}

	b, err := json.MarshalIndent(x, "", "  ")
	if err != nil {
		log.Fatal("\nJSON marshaling error: ", err)
	}
	fmt.Println(string(b))
}

// NormalizePath @todo rename to ToSlash
func NormalizePath(path string) string {
	return filepath.ToSlash(path)
}

// Pause @todo
func Pause() {
	fmt.Print("Press 'Enter' to continue: ")
	bufio.NewReader(os.Stdin).ReadBytes('\n')
	fmt.Printf("\n")
}

// TimespecToTime @todo
func TimespecToTime(ts syscall.Timespec) time.Time {
	return time.Unix(int64(ts.Sec), int64(ts.Nsec))
}

// TrimSuffix @todo
func TrimSuffix(s, suffix string) string {
	if strings.HasSuffix(s, suffix) {
		s = s[:len(s)-len(suffix)]
	}
	return s
}
