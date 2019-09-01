package controllers

import (
	"encoding/json"
	"net/http"
	"resumator-backend/logger"

	"go.uber.org/zap"
)

// JSONResponse ...
type JSONResponse struct {
	Message string `json:"message,omitempty"`
}

var log *zap.SugaredLogger

// AuthController ...
func AuthController(writer http.ResponseWriter, request *http.Request) {
	log = logger.GetLogger()
	params := request.URL.Query()
	writer.Header().Set("Content-Type", "application/json")
	if params["code"] == nil {
		log.Infow("error occurred, request params empty")
		writer.WriteHeader(http.StatusBadRequest)
		badRequestResponse, err := json.Marshal(JSONResponse{
			Message: "Bad Request",
		})
		if err != nil {
			writer.WriteHeader(http.StatusInternalServerError)
			return
		}
		writer.Write(badRequestResponse)
		return
	}

	log.Infow("logging", "params", params)
	writer.WriteHeader(http.StatusOK)
}
