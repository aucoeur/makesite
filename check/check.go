package check

import (
	"fmt"
	"os"
	"path/filepath"
)

// Observe checks for error and panics if there is one.  Thanks gobyexample.com
func Observe(e error) {
	if e != nil {
		panic(e)
	}
}

// DoesMatch checks if the path passed matches the flag
func DoesMatch(flag, path string) bool {

	info, err := os.Stat(path)
	Observe(err)
	switch {
	case flag == "file" && info.IsDir():
		fmt.Println("This appears to be a directory.  Please use -dir for directory, otherwise please choose a single file")
		return false
	case flag == "dir" && !info.IsDir():
		fmt.Println("This appears to be a file.  Please use -file for single file, otherwise please choose a directory")
		return false
	default:
		return true
	}

}

// GetFilesInDir checks for all files in a directory including sub-directories and returns an array with the files
func GetFilesInDir(dir string) []string {
	files := make([]string, 0)

	// Only add to files array if path is not a directory
	err := filepath.Walk(dir, func(path string, info os.FileInfo, e error) error {
		if !info.IsDir() {
			files = append(files, path)
		}
		return e
	})
	Observe(err)
	return files
}
