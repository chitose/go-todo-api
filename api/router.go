package api

import (
	"github.com/chitose/todo-api/api/auth"
	"github.com/chitose/todo-api/api/home"
	"github.com/chitose/todo-api/api/project"
	"github.com/gorilla/mux"
)

func SetupRouter(r *mux.Router) {
	auth.ConfigureRoute(r)
	home.ConfigureRouter(r)
	project.ConfigureRouter(r)
}
