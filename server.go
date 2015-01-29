package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
)

func Home(w http.ResponseWriter, req *http.Request) {
	render(w, "index.html")
}

func About() {
	render(w, "about.html")
}

func render(w http.ResponseWriter, tmpl string) {
	tmpl = fmt.Sprintf("template/%s", tmpl)

	t, err := template.ParseFiles(tmpl)
	if err != nil {
		log.Print("Template parsing error: ", err)
	}

	err = t.Execute(w, "")
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
	err := http.ListenAndServe(":8000", nil)

	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
