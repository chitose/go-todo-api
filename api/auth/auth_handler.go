package auth

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/markbates/goth/gothic"
)

func authHandler(res http.ResponseWriter, req *http.Request) {
	gothic.BeginAuthHandler(res, req)
}

func ConfigureRoute(r *mux.Router) {
	r.HandleFunc("/auth/{provider}", authHandler)
	r.HandleFunc("/auth/{provider}/callback", authCallbackHandler)
}
