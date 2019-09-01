package routes

import (
	"net/http"
	"resumator-backend/controllers"
	services "resumator-backend/services/impl"

	"github.com/gorilla/mux"
)

func InitializeRouter() *mux.Router {
	router := mux.NewRouter()
	router.HandleFunc("/auth/accessToken", getAuthController().GetAccessToken)
	http.Handle("/", router)
	return router
}

func getAuthController() *controllers.AuthController {
	controller := &controllers.AuthController{
		AuthService: getLinkedInAuthService(),
	}
	return controller
}

func getLinkedInAuthService() *services.LinkedInAuthService {
	service := &services.LinkedInAuthService{
		Client: &http.Client{},
	}
	return service
}
