package util

import (
	"time"
	"fmt"
	"net/http"
	"path/filepath"
	"strings"
	"os"
	"path"
	"errors"
)

// util from github.com/dyxj/gomod util.go

// A DirD implements FileSystem using the native file system restricted to a
// specific directory tree.
type DirD struct {
	Dir string
	Def string
}

// Open : if a file is not found in DirD.Dir,
// it will serve a default file specified by DirD.Def
// While the FileSystem.Open method takes '/'-separated paths, a Dir's string
// value is a filename on the native file system, not a URL, so it is separated
// by filepath.Separator, which isn't necessarily '/'.
//
// An empty Dir is treated as ".".
func (d DirD) Open(name string) (http.File, error) {
	if filepath.Separator != '/' && strings.ContainsRune(name, filepath.Separator) ||
		strings.Contains(name, "\x00") {
		return nil, errors.New("http: invalid character in file path")
	}
	dir := string(d.Dir)
	if dir == "" {
		dir = "."
	}
	f, err := os.Open(filepath.Join(dir, filepath.FromSlash(path.Clean("/"+name))))
	if err != nil {
		fDef, errDef := os.Open(filepath.Join(dir, filepath.FromSlash(path.Clean("/"+d.Def))))
		if errDef != nil {
			return nil, errDef
		}
		return fDef, nil
	}
	return f, nil
}

// TrackTime : Used to track execution time,
// Can be used with defer to get execution time of function
// ie:
// defer log.Println(TrackTime(time.Now(), "funcName()"))
func TrackTime(start time.Time, fname string) string {
	//t := time.Since(start).Nanoseconds()
	//return fmt.Sprintf("%s execution time: %s\n", fname, strconv.Itoa(int(t)))
	t := time.Since(start)
	return fmt.Sprintf("%s execution time: %s\n", fname, t)
}