package util

import (
	"bufio"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"runtime"
	"strings"
	"syscall"
	"time"
)

// Basename @todo
// basename removes trailing slashes and the leading
// directory name and drive letter from path name.
func Basename(name string) string {
	// Remove drive letter
	if len(name) == 2 && name[1] == ':' {
		name = "."
	} else if len(name) > 2 && name[1] == ':' {
		name = name[2:]
	}
	i := len(name) - 1
	// Remove trailing slashes
	for ; i > 0 && (name[i] == '/' || name[i] == '\\'); i-- {
		name = name[:i]
	}
	// Remove leading directory name
	for i--; i >= 0; i-- {
		if name[i] == '/' || name[i] == '\\' {
			name = name[i+1:]
			break
		}
	}
	return name
}

// Dirname @todo
func Dirname(path string) string {
	dir, _ := filepath.Split(path)
	return TrimSuffix(TrimSuffix(dir, "/"), "\\")
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

// NormalizePath @todo
func NormalizePath(path string) string {
	if runtime.GOOS == "windows" {
		return strings.Replace(path, "\\", "/", -1)
	}
	return path
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
