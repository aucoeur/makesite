package pages

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
	"text/template"

	"github.com/aucoeur/makesite/check"
	"github.com/fatih/color"
	"github.com/russross/blackfriday/v2"
)

// Entry is blog post struct
type Entry struct {
	Title string
	Body  string
}

// CreateHTMLFromTemplate takes a path string of a text file along with the path of a template dir and creates a new HTML post from it
func CreateHTMLFromTemplate(file string, tmplDir string) {
	// Some filepath wrangling
	ext := filepath.Ext(file)
	base := filepath.Base(file)
	filename := strings.TrimSuffix(base, ext)

	// Reading
	fileContents, err := ioutil.ReadFile(file)
	check.Report(err)

	switch strings.ToUpper(ext) {
	case ".MD":
		fileContents = blackfriday.Run(fileContents)
		fallthrough
	case ".TXT":
		// Assumes first line contains title and separated by body contents by 1+ newline
		fileSplit := strings.SplitN(string(fileContents), "\n", 2)

		ent := Entry{
			Title: fileSplit[0],
			Body:  strings.TrimSpace(fileSplit[1]),
		}
		// Create new file
		postDir := "posts/"
		w, err := os.Create(filepath.Join(postDir, filename) + ".html")
		check.Report(err)

		// Template stuff
		t := template.Must(template.ParseFiles(filepath.Join(tmplDir, "template.tmpl"), filepath.Join(tmplDir, "header.tmpl")))
		t.ExecuteTemplate(w, "template.tmpl", ent)
		check.Report(err)

		fmt.Printf(filename + ".html saved successfully\n")
	}

}

// BatchCreateHTMLFromTemplate runs through files array iteratively and createsHTMLfromtemplate or HTML from MD
func BatchCreateHTMLFromTemplate(files []string, tmplDir string) {
	for _, file := range files {
		fmt.Printf("Creating html from: %s ...", file)
		CreateHTMLFromTemplate(file, tmplDir)
	}

	// Print to console with colors when finished
	greenbold := color.New(color.FgGreen, color.Bold).SprintFunc()
	bold := color.New(color.Bold).SprintFunc()
	fmt.Printf("%s %s %s", greenbold("Success! Generated"), bold(len(files)), greenbold("pages."))
}
