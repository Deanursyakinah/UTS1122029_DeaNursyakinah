package controller

import (
	"encoding/json"
	"net/http"

	m "UTS_1122029/model"
)

func SendErrorResponse(w http.ResponseWriter, kode int) {
	w.Header().Set("Content-Type", "application/json")
	var response m.ErrorResponse
	response.Status = kode //400 bad req, 404 not found, 500 internal server error, 401 unauthorized

	json.NewEncoder(w).Encode(response)
}

func SendSuccesResponse(w http.ResponseWriter, kode int) {
	w.Header().Set("Content-Type", "application/json")
	var response m.SuccessResponse
	response.Status = kode
	json.NewEncoder(w).Encode(response)
}
