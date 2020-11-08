package auth

import (
	"encoding/json"
	"net/http"

	"github.com/renosyah/go-postgre/api"
)

func RespondError(w http.ResponseWriter, message string, status int) {
	resp := api.Response{
		Status: status,
		Data:   nil,
		Errors: []api.Error{
			{
				Log:        "",
				StatusCode: status,
				Message:    message,
			},
		},
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	if err := json.NewEncoder(w).Encode(&resp); err != nil {
		return
	}
}
