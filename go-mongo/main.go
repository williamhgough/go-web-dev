package main

import (
	"github.com/julienschmidt/httprouter"
	"github.com/nytimes/gziphandler"
	"github.com/williamhgough/go-web-dev/go-mongo/controllers"
	"html/template"
	"net/http"
)

var tpl *template.Template

func main() {
	server := httprouter.New()
	uc := controllers.NewUserController()

	server.GET("/", index)
	server.GET("/user/:id", uc.GetUser)
	server.POST("/user", uc.CreateUser)
	server.DELETE("/user/:id", uc.DeleteUser)

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
