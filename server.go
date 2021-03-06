package main

import (
	"fmt"
	"html/template"
	"io"
	"log"
	"net/http"
	"time"
)

const STATIC_URL string = "/static/"
const STATIC_ROOT string = "static/"

// This is the struct made available to our templates
type Context struct {
	Title  string
	Static string
}

func Home(w http.ResponseWriter, req *http.Request) {
	context := Context{Title: "Welcome! Home Page"}
	render(w, "index.html")
}

func About() {
	context := Context{Title: "About Page"}
	render(w, "about.html")
}

func render(w http.ResponseWriter, tmpl string, context Context) {
	context.Static = STATIC_URL
	tmpl_list = []string{"templates/layout.html", fmt.Sprintf("template/%s", tmpl)}

	t, err := template.ParseFiles(tmpl_list...)
	if err != nil {
		log.Print("Template parsing error: ", err)
	}

	err = t.Execute(w, context)
	if err != nil {
		log.Print("Template execution erro: ", err)
	}
}

func StaticHandler(w http.ResponseWriter, req *http.Request) {
	static_file := req.URL.Path[len(STATIC_URL):]
	if len(static_file) != 0 {
		f, err := http.Dir(STATIC_ROOT).Open(static_file)
		if err == nil {
			content := io.ReadSeeker(f)
			http.ServeContent(w, req, static_file, time.Now(), content)
			return
		}
	}

	http.NotFound(w, req)
}

func main() {
	http.HandleFunc("/", Home)
	http.HandleFunc("/about/", About)
	http.HandleFunc(STATIC_URL, StaticHandler)

	err := http.ListenAndServe(":8000", nil)

	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
