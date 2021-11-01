package restapi

import (
	"net/http"

	"gitlab.com/ulexxander/remoconf/storage"
)

func (h *Handler) getProjectsAll(w http.ResponseWriter, r *http.Request) {
	res, err := h.projects.GetAll()
	if err != nil {
		h.resError(w, err)
		return
	}
	h.resData(w, res)
}

func (h *Handler) postProject(w http.ResponseWriter, r *http.Request) {
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
