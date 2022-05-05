package home

import (
	"html/template"
	"net/http"

	"github.com/gorilla/mux"
)

func homeHandler(res http.ResponseWriter, req *http.Request) {
	t, _ := template.ParseFiles("templates/index.html")
	t.Execute(res, false)
}

func ConfigureRouter(r *mux.Router) {
	r.HandleFunc("/", homeHandler)
}
