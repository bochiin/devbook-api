package routes

import (
	"api/src/controllers"
	"net/http"
)

var rotaLogin = Route{
	Uri:     "/login",
	Method:  http.MethodPost,
	Func:    controllers.Login,
	HasAuth: false,
}
