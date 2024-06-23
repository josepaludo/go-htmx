package server

import (
	routes "josepaludo/go-htmx/server/routes"
	"net/http"
)

type Route struct {
	path       string
	controller func(responseWriter http.ResponseWriter, request *http.Request)
}

var ROUTES = []Route{
	{"GET /", routes.Index},
}
