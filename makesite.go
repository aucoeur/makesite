package main

import (
	"flag"
	"fmt"
	"html/template"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"

	"github.com/fatih/color"
)

// Entry is blog post struct
type Entry struct {
	Title string
	Body  string
}

// Thanks gobyexample.com
func check(e error) {
	if e != nil {
		panic(e)
	}
}

func createHTMLFromTemplate(file string, tmpl string) {
	// Some filepath wrangling
	ext := filepath.Ext(file)
	base := filepath.Base(file)
	filename := strings.TrimSuffix(base, ext)

	// Reading
	fileContents, err := ioutil.ReadFile(file)
	check(err)

	// Assumes first line contains title and separated by body contents by 1+ newline
	fileSplit := strings.SplitN(string(fileContents), "\n", 2)

	ent := Entry{
		Title: fileSplit[0],
		Body:  strings.TrimSpace(fileSplit[1]),
	}

	// Create new file
	postDir := "posts/"
	w, err := os.Create(filepath.Join(postDir, filename) + ".html")
	check(err)

	// Template stuff
	tmplDir := tmpl
	t := template.Must(template.ParseFiles(filepath.Join(tmplDir, "template.tmpl"), filepath.Join(tmplDir, "header.tmpl")))
	t.ExecuteTemplate(w, "template.tmpl", ent)
	check(err)

	fmt.Println(filename + ".html saved successfully")
}

func main() {

	// flags := flag.FlagSet()
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
	case file != "":
		info, err := os.Stat(file)
		check(err)
		if info.IsDir() {
			fmt.Println("This appears to be a directory.  Please use -dir for directory, otherwise please choose a single file")
		} else {
			fmt.Printf("file: ")
			createHTMLFromTemplate(file, tmplDir)
		}
		// If -dir is specified, check if file or directory before recursively looking for files
	case dir != "":
		dirInfo, err := os.Stat(dir)
		check(err)

		if dirInfo.IsDir() {
			files := make([]string, 0)

			// Only add to files array if path is not a directory
			err := filepath.Walk(dir, func(path string, info os.FileInfo, e error) error {
				if !info.IsDir() {
					files = append(files, path)
				}
				return e
			})
			check(err)

			// Loop through files and create for each file in array files
			for _, file := range files {
				fmt.Println("Creating html from: ", file)
				createHTMLFromTemplate(file, tmplDir)
			}

			// Print to console with colors when finished
			greenbold := color.New(color.FgGreen, color.Bold).SprintFunc()
			bold := color.New(color.Bold).SprintFunc()
			fmt.Printf("%s %s %s", greenbold("Success! Generated"), bold(len(files)), greenbold("pages."))

		} else {
			fmt.Println("This appears to be a file.  Please use -file for single file, otherwise please choose a directory")
		}

	}

}
