package routes

import (
	"net/http"

	"github.com/gorilla/mux"
)

type Route struct {
	Uri     string
	Method  string
	Func    func(http.ResponseWriter, *http.Request)
	HasAuth bool
}

func ConfigRoutes(r *mux.Router) *mux.Router {

	routes := userRoutes

	routes = append(routes, rotaLogin)

	for _, route := range routes {
		r.HandleFunc(route.Uri, route.Func).Methods(route.Method)
	}

	return r
}
