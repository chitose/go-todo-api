package main

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

const userApi = "/api/user"

func setupUserRouter(r *mux.Router) {
	r.HandleFunc(userApi, userIndexHandler)
}

func userIndexHandler(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("content-type", "application/json")
	res.WriteHeader(http.StatusOK)
	type ColorGroup struct {
		ID     int
		Name   string
		Colors []string
	}
	group := ColorGroup{
		ID:     1,
		Name:   "Reds",
		Colors: []string{"Crimson", "Red", "Ruby", "Maroon"},
	}
	b, _ := json.Marshal(group)
	res.Write(b)
}
