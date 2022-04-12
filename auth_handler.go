package main

import (
	"net/http"

	"github.com/markbates/goth/gothic"
)

func authHandler(res http.ResponseWriter, req *http.Request) {
	gothic.BeginAuthHandler(res, req)
}
