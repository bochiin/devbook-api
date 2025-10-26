package router

import (
	"api/src/router/routes"

	"github.com/gorilla/mux"
)

func CreateRouter() *mux.Router {
	return routes.ConfigRoutes(mux.NewRouter())
}
