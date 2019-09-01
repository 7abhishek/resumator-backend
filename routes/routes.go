package routes

import (
	"net/http"
	"resumator-backend/controllers"

	"github.com/gorilla/mux"
)

func InitializeRouter() *mux.Router {
	router := mux.NewRouter()
	router.HandleFunc("/auth", controllers.AuthController)
	http.Handle("/", router)
	return router
}
