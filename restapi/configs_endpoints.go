package restapi

import (
	"net/http"

	"gitlab.com/ulexxander/remoconf/storage"
)

// @ID GetConfigsAll
// @Produce json
// @Success 200 {object} []storage.Config
// @Failure default {object} ResponseError
// @Router /configs [get]
func (e *Endpoints) GetConfigsAll(w http.ResponseWriter, r *http.Request) {
	res, err := e.Configs.GetAll()
	if err != nil {
		e.resError(w, err)
		return
	}
	e.resData(w, res)
}

// @ID PostConfig
// @Accept json
// @Produce json
// @Param params body storage.ConfigCreateParams true "Config Create Params"
// @Success 200 {object} storage.CreatedItem
// @Failure default {object} ResponseError
// @Router /configs [post]
func (e *Endpoints) PostConfig(w http.ResponseWriter, r *http.Request) {
	var p storage.ConfigCreateParams
	if err := bodyJSON(r, &p); err != nil {
		e.resError(w, err)
		return
	}
	res, err := e.Configs.Create(p)
	if err != nil {
		e.resError(w, err)
		return
	}
	e.resData(w, res)
}
