package main

import (
	"encoding/json"
	"errors"
	"io"
	"net/http"
)

type JSONResponse struct {
	Error   bool   `json:"error"`
	Message string `json:"message"`
	Data    any    `json:"data,omitempty"`
}

func (app *Config) readJSON(response http.ResponseWriter, request *http.Request, data any) error {
	maxBytes := 1048576 //One mb
	request.Body = http.MaxBytesReader(response, request.Body, int64(maxBytes))

	decoder := json.NewDecoder(request.Body)
	err := decoder.Decode(data)
	if err != nil {
		return err
	}

	err = decoder.Decode(&struct{}{})
	if err != io.EOF {
		return errors.New("body must have only a single JSON value")
	}

	return nil
}

func (app *Config) writeJSON(response http.ResponseWriter, status int, data any, headers ...http.Header) error {
	out, err := json.Marshal(data)
	if err != nil {
		return err
	}

	if len(headers) > 0 {
		for key, value := range headers[0] {
			response.Header()[key] = value
		}
	}

	response.Header().Set("Content-Type", "application/json")
	response.WriteHeader(status)
	_, err = response.Write(out)
	if err != nil {
		return err
	}

	return nil
}

func (app *Config) errorJSON(response http.ResponseWriter, err error, status ...int) error {
	statusCode := http.StatusBadRequest
	if len(status) > 0 {
		statusCode = status[0]
	}

	var payload JSONResponse
	payload.Error = true
	payload.Message = err.Error()

	return app.writeJSON(response, statusCode, payload)
}
