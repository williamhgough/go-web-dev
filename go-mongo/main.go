package main

import (
	"github.com/julienschmidt/httprouter"
	"github.com/nytimes/gziphandler"
	"html/template"
	"net/http"
)

var tpl *template.Template

func main() {
	server := httprouter.New()
	server.GET("/", index)
	zipped := gziphandler.GzipHandler(server)
	http.ListenAndServe("localhost:8080", zipped)
}

func init() {
	tpl = template.Must(template.ParseGlob("templates/*.gohtml"))
}

func index(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	tpl.ExecuteTemplate(w, "index.gohtml", nil)
}
