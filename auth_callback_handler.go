package main

import (
	"fmt"
	"html/template"
	"net/http"
	"os"

	"github.com/golang-jwt/jwt"
	"github.com/markbates/goth"
	"github.com/markbates/goth/gothic"
)

type OauthData struct {
	User goth.User
	Jwt  string
}

func authCallbackHandler(res http.ResponseWriter, req *http.Request) {
	user, err := gothic.CompleteUserAuth(res, req)
	if err != nil {
		fmt.Fprintln(res, err)
		return
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256,
		jwt.MapClaims{
			"id": user.UserID,
		})

	hmacSecret := []byte(os.Getenv("APP_SECRET"))

	tokenString, err := token.SignedString(hmacSecret)

	if err != nil {
		fmt.Fprintln(res, err)
		return
	}

	t, _ := template.ParseFiles("templates/success.html")

	data := OauthData{user, tokenString}

	t.Execute(res, data)
}
