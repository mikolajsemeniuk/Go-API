package main

import (
	"encoding/json"
	"net/http"
)

func (app *application) WriteJson(w http.ResponseWriter, status int, data interface{}, wrap string) error {
	wrapper := make(map[string]interface{})
	wrapper[wrap] = data

	js, error := json.Marshal(wrapper)
	if error != nil {
		return error
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	w.Write(js)
	return nil
}

func (app *application) ErrorJSON(w http.ResponseWriter, err error) {
	_error := struct {
		Message string `json:"message"`
	}{
		Message: err.Error(),
	}

	app.WriteJson(w, http.StatusBadRequest, _error, "error")
}
