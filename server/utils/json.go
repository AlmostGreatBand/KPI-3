package utils

import (
	"encoding/json"
	"log"
	"net/http"
)

type errorMessage struct {
	ErrorMessage string `json:"errorMessage"`
}

func ResponseOk(rw http.ResponseWriter, res interface{}) {
	writeJson(rw, http.StatusOK, res)
}

/*
	for a successful PUT of an update to and existing resource no body needed,
	so status code 204 - No Content is more appropriate that 200 OK
 */
func ResponseOkNoBody(rw http.ResponseWriter) {
	writeJson(rw, http.StatusNoContent, nil)
}

func ResponseBadRequest(rw http.ResponseWriter, message string) {
	writeJson(rw, http.StatusBadRequest, errorMessage{ErrorMessage: message})
}

func ResponseInternalError(rw http.ResponseWriter, message string) {
	writeJson(rw, http.StatusInternalServerError, errorMessage{ErrorMessage: message})
}

func writeJson(rw http.ResponseWriter, status int, res interface{}) {
	rw.Header().Set("content-type", "application/json")
	rw.WriteHeader(status)
	if res != nil {
		if err := json.NewEncoder(rw).Encode(res); err != nil {
			log.Printf("error: something went wrong while encoding json")
		}
	}
}
