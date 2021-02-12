package main

import (
	"flag"
	"fmt"
	"html/template"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
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
	case flag.NFlag() == 0:
		fmt.Println("a flag is required")
		flag.PrintDefaults()
		os.Exit(1)
	case file != "":
		info, err := os.Stat(file)
		check(err)
		if info.IsDir() {
			fmt.Println("This appears to be a directory.  Please use -dir for directory, otherwise please choose a single file")
		} else {
			fmt.Printf("file: ")
			createHTMLFromTemplate(file, tmplDir)
		}
	case dir != "":
		dirInfo, err := os.Stat(dir)
		check(err)

		if dirInfo.IsDir() {
			files, err := ioutil.ReadDir(dir)
			check(err)
			for _, f := range files {
				fmt.Printf("file: %s", f.Name())
				createHTMLFromTemplate(dir+"/"+f.Name(), tmplDir)
			}
		} else {
			fmt.Println("This appears to be a file.  Please use -file for single file, otherwise please choose a directory")
		}

	}

}
