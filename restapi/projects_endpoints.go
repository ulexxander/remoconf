package restapi

import (
	"net/http"

	"gitlab.com/ulexxander/remoconf/storage"
)

// @Produce json
// @Success 200 {object} []storage.Project
// @Failure default {object} ResponseError
// @Router /projects [get]
func (e *Endpoints) GetProjectsAll(w http.ResponseWriter, r *http.Request) {
	res, err := e.Projects.GetAll()
	if err != nil {
		e.resError(w, err)
		return
	}
	e.resData(w, res)
}

// @Accept json
// @Produce json
// @Param params body storage.ProjectCreateParams true "Project Create Params"
// @Success 200 {object} storage.CreatedItem
// @Failure default {object} ResponseError
// @Router /projects [post]
func (e *Endpoints) PostProject(w http.ResponseWriter, r *http.Request) {
	var p storage.ProjectCreateParams
	if err := bodyJSON(r, &p); err != nil {
		e.resError(w, err)
		return
	}
	res, err := e.Projects.Create(p)
	if err != nil {
		e.resError(w, err)
		return
	}
	e.resData(w, res)
}
