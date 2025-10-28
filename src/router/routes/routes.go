package routes

import (
	"api/src/middlewares"
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

		if route.HasAuth {
			r.HandleFunc(
				route.Uri,
				middlewares.Logger(middlewares.Authenticate(route.Func)),
			).Methods(route.Method)
			continue
		}

		r.HandleFunc(route.Uri, middlewares.Logger(route.Func)).Methods(route.Method)
	}

	return r
}
