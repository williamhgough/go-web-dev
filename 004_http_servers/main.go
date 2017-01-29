package main

import (
	"net/http"
	"github.com/nytimes/gziphandler"
	"log"
	"net/url"
	"html/template"
)

type handle int

func (m handle) ServeHTTP(res http.ResponseWriter, req *http.Request) {
	err := req.ParseForm()
	if err != nil {
		log.Println(err)
	}

	data := struct {
		Submissions url.Values
	}{
		req.Form,
	}

	tpl.ExecuteTemplate(res, "index.gohtml", data)
}

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseFiles("templates/index.gohtml"))
}

func main() {
	var d handle
	withGz := gziphandler.GzipHandler(d)
	http.ListenAndServe(":8080", withGz)
}
