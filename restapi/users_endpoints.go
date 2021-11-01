package restapi

import (
	"net/http"

	"gitlab.com/ulexxander/remoconf/storage"
)

// @ID GetUserByID
// @Produce  json
// @Param id path int true "User ID"
// @Success 200 {object} storage.User
// @Failure default {object} ResponseError
// @Router /users/{id} [get]
func (h *Handler) getUserByID(w http.ResponseWriter, r *http.Request) {
	id, err := urlParamInt(r, "id")
	if err != nil {
		h.resError(w, err)
		return
	}
	res, err := h.users.GetByID(id)
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
	res, err := h.users.Create(p)
	if err != nil {
		h.resError(w, err)
		return
	}
	h.resData(w, res)
}
