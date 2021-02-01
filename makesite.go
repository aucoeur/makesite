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

// 	[TODO] Split into Title & Body
// func createEntry(content string) {
// 	fileSplit := strings.SplitN(string(content), "\n", 2)
// 	e := &Entry{
// 		Title: fileSplit[0],
// 		Body:  fileSplit[1],
// 	}
// 	fmt.Printf("%+v", e)
// 	return e
// }

func main() {

	// Some sandbox set up
	filePtr := flag.String("file", "", "a text file")
	flag.Parse()
	// file := "first-post.txt"
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
