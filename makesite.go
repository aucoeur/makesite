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

	fmt.Printf(filename + ".html saved successfully")
}

func main() {

	// Flags
	filePtr := flag.String("file", "sample.txt", "a text file to convert")
	// dirPtr := flag.String("dir", ".", "a directory of text files to convert")
	tDirPtr := flag.String("templates", "./templates", "a directory of templates")

	flag.Parse()

	// Point parameters to flags
	files := *filePtr
	// txtDir := *dirPtr
	tmplDir := *tDirPtr
	info, err := os.Stat(files)
	check(err)
	if info.IsDir() {

		files, err := ioutil.ReadDir(files)
		check(err)
		for _, file := range files {
			fmt.Printf("%s \n", file.Name())
			// createHTMLFromTemplate(filepath.Abs(file), tmplDir)
		}
	} else {
		fmt.Printf("file: ")
		createHTMLFromTemplate(files, tmplDir)
	}

}
