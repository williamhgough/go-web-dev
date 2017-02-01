package main

import (
	"github.com/julienschmidt/httprouter"
	"github.com/nytimes/gziphandler"
	"github.com/williamhgough/go-web-dev/go-mongo/controllers"
	"gopkg.in/mgo.v2"
	"html/template"
	"log"
	"net/http"
)

var tpl *template.Template

func main() {
	server := httprouter.New()
	uc := controllers.NewUserController(getSession())

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

func getSession() *mgo.Session {
	// Connect to our local storage
	s, err := mgo.Dial("mongodb://admin:will2309@ds135069.mlab.com:35069/devtheweb")
	// Check if connection error
	if err != nil {
		log.Fatalln(err)
	}
	return s
}
