package project

import (
	"net/http"

	"github.com/chitose/todo-api/api/base"
	"github.com/chitose/todo-api/schema/dao"
	"github.com/gorilla/mux"
)

func ConfigureRouter(router *mux.Router) {
	router.HandleFunc("/api/projects", getAllProject)
}

// GetAllProject is a function to get a slice of record(s) from project table in the main database
// @Summary Get list of Project
// @Tags Project
// @Description GetAllProject is a handler to get a slice of record(s) from project table in the main database
// @Accept  json
// @Produce  json
// @Param   page     query    int     false        "page requested (defaults to 0)"
// @Param   pagesize query    int     false        "number of records in a page  (defaults to 20)"
// @Param   order    query    string  false        "db sort order column"
// @Success 200 {object} api.PagedResults{data=[]model.Project}
// @Failure 400 {object} api.HTTPError
// @Failure 404 {object} api.HTTPError
// @Router /project [get]
// http "http://localhost:8080/project?page=0&pagesize=20" X-Api-User:user123
func getAllProject(w http.ResponseWriter, r *http.Request) {
	ctx := base.InitializeContext(r)

	page, err := base.ReadInt(r, "page", 0)

	if err != nil || page < 0 {
		base.ReturnError(ctx, w, r, dao.ErrBadParams)
		return
	}

	pagesize, err := base.ReadInt(r, "pagesize", 20)

	order := r.FormValue("order")

	records, totalRows, err := dao.GetAllProject(ctx, page, pagesize, order)
	if err != nil {
		base.ReturnError(ctx, w, r, err)
		return
	}

	result := &base.PagedResults{Page: page, PageSize: pagesize, Data: records, TotalRecords: totalRows}
	base.WriteJSON(ctx, w, result)
}
