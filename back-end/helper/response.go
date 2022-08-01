package helper

import (
	"JUALiND/schema"
	"encoding/json"
	"net/http"
)

func SuccessResponseJSON(w http.ResponseWriter, msg string, body interface{}) {

	response := schema.JSONResponse{
		Message: msg,
		Body:    body,
	}

	responseByte, _ := json.Marshal(response)
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(responseByte)
}

func ErrorResponseJSON(w http.ResponseWriter, err error, msg string, statusCode int) {
	response := schema.JSONResponse{
		Message: msg,
		Error:   err.Error(),
	}

	responseByte, _ := json.Marshal(response)

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	w.Write(responseByte)
}
