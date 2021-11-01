package restapi

import (
	"net/http"

	"gitlab.com/ulexxander/remoconf/service/configs"
)

// @ID GetConfigsByProject
// @Produce json
// @Param id path int true "User ID"
// @Success 200 {object} []storage.Config
// @Failure default {object} ResponseError
// @Router /projects/{id}/configs [get]
func (e *Endpoints) GetConfigsByProject(w http.ResponseWriter, r *http.Request) {
	id, err := urlParamInt(r, "id")
	if err != nil {
		e.resError(w, err)
		return
	}
	res, err := e.Configs.GetByProject(id)
	if err != nil {
		e.resError(w, err)
		return
	}
	e.resData(w, res)
}

// @ID PostConfig
// @Accept json
// @Produce json
// @Param params body configs.ConfigCreateParams true "Config Create Params"
// @Success 200 {object} storage.CreatedItem
// @Failure default {object} ResponseError
// @Router /configs [post]
func (e *Endpoints) PostConfig(w http.ResponseWriter, r *http.Request) {
	var p configs.ConfigCreateParams
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
