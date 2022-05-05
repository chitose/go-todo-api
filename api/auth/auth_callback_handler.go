package auth

import (
	"fmt"
	"html/template"
	"net/http"
	"os"

	"github.com/chitose/todo-api/schema/dao"
	"github.com/chitose/todo-api/schema/model"
	"github.com/golang-jwt/jwt"
	"github.com/guregu/null"
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

	ctx := req.Context()

	dbUser, err := dao.GetUser(ctx, user.UserID)
	if dbUser.ID == "" || err != nil {
		// add the authenticated user
		dao.AddUser(ctx, &model.User{
			ID:          user.UserID,
			DisplayName: null.StringFrom(user.FirstName + " " + user.LastName),
			Photo:       null.StringFrom(user.AvatarURL),
			Email:       null.StringFrom(user.Email),
		})

		// and default project (inbox)
		dao.AddProject(ctx, &model.Project{
			DefaultInbox: 1,
			Name:         "Inbox",
			View:         1,
			Archived:     0,
		})
	}

	t.Execute(res, data)
}
