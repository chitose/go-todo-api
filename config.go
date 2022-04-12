package main

import (
	"os"

	"github.com/gorilla/sessions"
	"github.com/markbates/goth"
	"github.com/markbates/goth/gothic"
)

func setup() {
	maxAge := 86400 * 30 // 30 days
	isProd := false

	store := sessions.NewCookieStore([]byte(os.Getenv("APP_SECRET")))
	store.MaxAge(maxAge)
	store.Options.Path = "/"
	store.Options.HttpOnly = true // HttpOnly should always be enabled
	store.Options.Secure = isProd

	gothic.Store = store

	callBackUrlPrefix := "http://localhost:3000/auth"

	goth.UseProviders(
		getGoogleProvider(callBackUrlPrefix),
	)
}
