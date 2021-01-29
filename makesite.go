package main

import (
	"fmt"
	"html/template"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
)

// type Entry struct {
// 	Title string
// 	Body  string
// }

// Thanks gobyexample.com
func check(e error) {
	if e != nil {
		panic(e)
	}
}

// 	[TODO] Split into Title & Body
// func getTitle(file string) {
// 	f, err := os.Open(file)
// 	defer f.Close()
// 	check(err)

// 	s := bufio.NewScanner(f)
// 	entry := new(Entry)

// 	for s.Scan() {
// 		if entry.Title == "" {
// 			entry.Title = s.Text()
// 		} else {

// 		}
// 	}
// 	check(err)
// }

func main() {

	// Some sandbox set up
	file := "first-post.txt"
	ext := filepath.Ext(file)
	filename := strings.TrimSuffix(file, ext)

	// Reading
	fileContents, err := ioutil.ReadFile(file)
	check(err)

	// Create new file
	w, err := os.Create(filename + ".html")
	check(err)

	// Template stuff
	t := template.Must(template.ParseFiles("template.tmpl"))
	t.ExecuteTemplate(w, "template.tmpl", string(fileContents))

	fmt.Printf(filename + ".html saved successfully")
}
