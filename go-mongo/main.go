package main

import (
	"encoding/json"
	"fmt"
	"github.com/julienschmidt/httprouter"
	"github.com/nytimes/gziphandler"
	"github.com/williamhgough/go-web-dev/go-mongo/models"
	"html/template"
	"net/http"
)

var tpl *template.Template

func main() {
	server := httprouter.New()
	server.GET("/", index)
	server.GET("/user/:id", getUser)
	zipped := gziphandler.GzipHandler(server)
	http.ListenAndServe("localhost:8080", zipped)
}

func init() {
	tpl = template.Must(template.ParseGlob("templates/*.gohtml"))
}

func index(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	w.Header().Set("Content-Type", "text/html; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	tpl.ExecuteTemplate(w, "index.gohtml", nil)
}

func getUser(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	u := models.User{
		Name:   "William Gough",
		Gender: "male",
		Age:    23,
		Id:     p.ByName("id"),
	}

	// Marshal into Json
	uj, _ := json.Marshal(u)
	// Write content type, status code
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "%s\n", uj)
}
