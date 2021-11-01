package restapi

import (
	"net/http"

	"gitlab.com/ulexxander/remoconf/storage"
)

func (h *Handler) getUserByID(w http.ResponseWriter, r *http.Request) {
	id, err := urlParamInt(r, "id")
	if err != nil {
		h.resError(w, err)
		return
	}
	res, err := h.usersStore.GetByID(id)
	if err != nil {
		h.resError(w, err)
		return
	}
	h.resData(w, res)
}

func (h *Handler) postUser(w http.ResponseWriter, r *http.Request) {
	var p storage.UserCreateParams
	if err := bodyJSON(r, &p); err != nil {
		h.resError(w, err)
		return
	}
	res, err := h.usersStore.Create(p)
	if err != nil {
		h.resError(w, err)
		return
	}
	h.resData(w, struct{ ID int }{ID: res})
}
