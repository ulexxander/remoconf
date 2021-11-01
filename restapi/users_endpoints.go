package restapi

import (
	"net/http"

	"gitlab.com/ulexxander/remoconf/storage"
)

// @Produce json
// @Param id path int true "User ID"
// @Success 200 {object} storage.User
// @Failure default {object} ResponseError
// @Router /users/{id} [get]
func (e *Endpoints) GetUserByID(w http.ResponseWriter, r *http.Request) {
	id, err := urlParamInt(r, "id")
	if err != nil {
		e.resError(w, err)
		return
	}
	res, err := e.Users.GetByID(id)
	if err != nil {
		e.resError(w, err)
		return
	}
	e.resData(w, res)
}

// @Accept json
// @Produce json
// @Param params body storage.UserCreateParams true "User Create Params"
// @Success 200 {object} storage.CreatedItem
// @Failure default {object} ResponseError
// @Router /users [post]
func (e *Endpoints) PostUser(w http.ResponseWriter, r *http.Request) {
	var p storage.UserCreateParams
	if err := bodyJSON(r, &p); err != nil {
		e.resError(w, err)
		return
	}
	res, err := e.Users.Create(p)
	if err != nil {
		e.resError(w, err)
		return
	}
	e.resData(w, res)
}
