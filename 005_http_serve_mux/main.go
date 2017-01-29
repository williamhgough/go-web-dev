package main

import (
	"io"
	"net/http"
)

func homePage(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "home page")
}

func aboutPage(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "about page")
}

func main() {
	http.HandleFunc("/home", homePage)
	http.HandleFunc("/about", aboutPage)
	http.HandleFunc("/*", http.NotFound)

	http.ListenAndServe(":8080", nil)
}
