package routes

import (
	"api/src/controllers"
	"net/http"
)

var userRoutes = []Route{
	{
		Uri:     "/user",
		Method:  http.MethodPost,
		Func:    controllers.CreateUser,
		HasAuth: false,
	},
	{
		Uri:     "/user",
		Method:  http.MethodGet,
		Func:    controllers.GetUsers,
		HasAuth: true,
	},
	{
		Uri:     "/user/{id}",
		Method:  http.MethodGet,
		Func:    controllers.GetUser,
		HasAuth: true,
	},
	{
		Uri:     "/user/{id}",
		Method:  http.MethodPut,
		Func:    controllers.UpdateUser,
		HasAuth: true,
	},
	{
		Uri:     "/user/{id}",
		Method:  http.MethodDelete,
		Func:    controllers.DeleteUser,
		HasAuth: true,
	},
}
