package controller

import (
	"log"
	"net/http"
)

type errorResponse struct {
	Message string `json:"message"`
}

func NewErrorResponse(w http.ResponseWriter, r *http.Request, statusCode int, message string) {
	log.Println("error:", message, "status code:", statusCode)
	http.Error(w, message, statusCode)
}
