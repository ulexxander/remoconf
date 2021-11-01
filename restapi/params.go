package restapi

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
	"github.com/pkg/errors"
)

func urlParamInt(r *http.Request, key string) (int, error) {
	val := chi.URLParam(r, key)
	valInt, err := strconv.Atoi(val)
	if err != nil {
		return 0, errors.Wrapf(err, "%s - invalid int", key)
	}
	return valInt, nil
}

func bodyJSON(r *http.Request, reqBody interface{}) error {
	err := json.NewDecoder(r.Body).Decode(&reqBody)
	if err != nil {
		return errors.Wrap(err, "decoding body json")
	}
	return nil
}
