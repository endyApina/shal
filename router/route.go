package router

import (
	"shal/controller"

	"github.com/gorilla/mux"
)

//SetupRouter function to setup all routing endpoints
func SetupRouter() *mux.Router {
	router := mux.NewRouter()
	router.HandleFunc("/", controller.SetupController)

	return router
}
