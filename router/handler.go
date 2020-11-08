package router

import (
	"encoding/json"
	"net/http"

	"github.com/renosyah/go-postgre/api"
)

type (
	HandlerFunc func(http.ResponseWriter, *http.Request) (interface{}, *api.Error)
)

func (fn HandlerFunc) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	errs := []api.Error{}
	status := http.StatusOK

	r.ParseForm()
	w.Header().Set("Content-Type", "application/json")

	data, err := fn(w, r)
	if err != nil {
		errs = append(errs, api.Error{
			Log:        err.Log,
			StatusCode: err.StatusCode,
			Message:    err.Message,
		})
		status = err.StatusCode
	}

	resp := api.Response{
		Status: status,
		Data:   data,
		Errors: errs,
	}
	if err := json.NewEncoder(w).Encode(&resp); err != nil {
		return
	}
}
