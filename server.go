package main

import (
	"io"
	"log"
	"net/http"
)

func Home(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "Hellooo Golang!")
}

func main() {
	http.HandleFunc("/", Home)
	err := http.ListenAndServe(":8000", nil)

	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
