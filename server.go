package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
)

func Home(w http.ResponseWriter, req *http.Request) {

}

func About() {

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

func main() {
	http.HandleFunc("/", Home)
	err := http.ListenAndServe(":8000", nil)

	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
