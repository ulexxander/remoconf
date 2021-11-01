package restapi

import (
	"net/http"

	"gitlab.com/ulexxander/remoconf/storage"
)

// @Produce json
// @Success 200 {object} []storage.Project
// @Failure default {object} ResponseError
// @Router /projects [get]
func (h *Handler) getProjectsAll(w http.ResponseWriter, r *http.Request) {
	res, err := h.projects.GetAll()
	if err != nil {
		h.resError(w, err)
		return
	}
	h.resData(w, res)
}

// @Accept json
// @Produce json
// @Param params body storage.ProjectCreateParams true "Project Create Params"
// @Success 200 {object} storage.CreatedItem
// @Failure default {object} ResponseError
// @Router /projects [post]
func (h *Handler) PostProject(w http.ResponseWriter, r *http.Request) {
	var p storage.ProjectCreateParams
	if err := bodyJSON(r, &p); err != nil {
		h.resError(w, err)
		return
	}
	res, err := h.projects.Create(p)
	if err != nil {
		h.resError(w, err)
		return
	}
	h.resData(w, res)
}
