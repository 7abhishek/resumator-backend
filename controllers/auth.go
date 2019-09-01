package controllers

import (
	"encoding/json"
	"net/http"
	"resumator-backend/logger"
	"resumator-backend/services"

	"go.uber.org/zap"
)

// JSONResponse ...
type JSONResponse struct {
	Message string `json:"message,omitempty"`
}

var log *zap.SugaredLogger

type AuthController struct {
	AuthService services.AuthService
}

// AuthController ...
func (controller *AuthController) GetAccessToken(writer http.ResponseWriter, request *http.Request) {

	writer.Header().Set("Access-Control-Allow-Origin", "*")
	log = logger.GetLogger()
	params := request.URL.Query()
	writer.Header().Set("Content-Type", "application/json")
	code := params["code"]
	if code == nil {
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

	log.Infow("logging", "params", code)
	accessToken, err := controller.AuthService.GetAccessToken(code[0])
	if err != nil {
		writer.WriteHeader(http.StatusInternalServerError)
		return
	}

	response, err := json.Marshal(accessToken)
	if err != nil {
		writer.WriteHeader(http.StatusInternalServerError)
		return
	}
	log.Infof("controller accessToken %s , %v", string(response), accessToken)
	writer.WriteHeader(http.StatusOK)
	writer.Write(response)
}
