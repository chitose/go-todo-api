package main

import (
	"html/template"
	"net/http"
)

func homeHandler(res http.ResponseWriter, req *http.Request) {
	t, _ := template.ParseFiles("templates/index.html")
	t.Execute(res, false)
}
