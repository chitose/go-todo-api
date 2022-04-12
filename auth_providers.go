package main

import (
	"os"

	"github.com/markbates/goth"
	"github.com/markbates/goth/providers/google"
)

func getGoogleProvider(callbackPrefix string) goth.Provider {
	return google.New(os.Getenv("GOOGLE_CLIENT_ID"),
		os.Getenv("GOOGLE_CLIENT_SECRET"),
		callbackPrefix+"/google/callback", "email", "profile")
}
