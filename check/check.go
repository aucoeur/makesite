package check

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

// Report checks for error and panics if there is one.  Thanks gobyexample.com
func Report(e error) {
	if e != nil {
		panic(e)
	}
}

// DoesMatch checks if the path passed matches the flag
func DoesMatch(flag, path string) bool {

	info, err := os.Stat(path)
	Report(err)
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
		if !info.IsDir() && isTextFile(path) {
			files = append(files, path)
		}
		return e
	})
	Report(err)
	return files
}

func isTextFile(path string) bool {
	ext := strings.ToUpper(filepath.Ext(path))
	if ext == ".TXT" || ext == ".MD" {
		fmt.Printf("found %s\n", path)
		return true
	}
	// fmt.Printf("skip %s\n", path)
	return false
}
