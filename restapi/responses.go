package restapi

import (
	"encoding/json"
	"net/http"
)

type ResponseData struct {
	Data interface{}
}

type ResponseError struct {
	Error string
}

func (e *Endpoints) resData(w http.ResponseWriter, data interface{}) {
	e.resJSON(w, ResponseData{
		Data: data,
	})
}

func (e *Endpoints) resError(w http.ResponseWriter, err error) {
	e.resJSON(w, ResponseError{
		Error: err.Error(),
	})
}

func (e *Endpoints) resJSON(w http.ResponseWriter, v interface{}) {
	err := json.NewEncoder(w).Encode(v)
	if err != nil {
		e.Logger.Printf("failed to write response json: %s", err)
		httpError(w, http.StatusInternalServerError)
	}
}

// httpError is like http.Error
// but it always sets text to status text of the specified code
func httpError(w http.ResponseWriter, code int) {
	http.Error(w, http.StatusText(code), code)
}
