package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/aucoeur/makesite/check"
	"github.com/aucoeur/makesite/pages"
)

func main() {

	// Flags
	filePtr := flag.String("file", "", "a text file to convert")
	dirPtr := flag.String("dir", "", "a directory of text files to convert")
	tmplDirPtr := flag.String("templates", "./templates", "a directory of templates")

	flag.Parse()

	// Point parameters to flags
	file := *filePtr
	dir := *dirPtr
	tmplDir := *tmplDirPtr

	switch {
	// If no flags are provided, show available flags and exit
	case flag.NFlag() == 0:
		fmt.Println("a flag is required")
		flag.PrintDefaults()
		os.Exit(1)

	// If -file is specified, check if file or directory before attempting to create HTML
	case file != "" && check.DoesMatch("file", file):
		pages.CreateHTMLFromTemplate(file, tmplDir)

	// If -dir is specified, check if file or directory before recursively looking for files
	case dir != "" && check.DoesMatch("dir", dir):
		files := check.GetFilesInDir(dir)
		pages.BatchCreateHTMLFromTemplate(files, tmplDir)

	}

}
