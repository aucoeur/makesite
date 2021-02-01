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

func main() {

	// Some sandbox set up
	// file := "first-post.txt"
	filePtr := flag.String("file", "sample.txt", "a text file")
	flag.Parse()
	file := *filePtr

	ext := filepath.Ext(file)
	filename := strings.TrimSuffix(file, ext)

	// Reading
	fileContents, err := ioutil.ReadFile(file)
	check(err)

	fileSplit := strings.SplitN(string(fileContents), "\n", 2)
	e := Entry{
		Title: fileSplit[0],
		Body:  strings.TrimSpace(fileSplit[1]),
	}
	// Create new file
	w, err := os.Create(filename + ".html")
	check(err)

	// Template stuff
	t := template.Must(template.ParseFiles("templates/template.tmpl", "templates/header.tmpl"))
	t.ExecuteTemplate(w, "template.tmpl", e)
	check(err)

	fmt.Printf(filename + ".html saved successfully")
}
