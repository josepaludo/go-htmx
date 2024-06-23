package server

import "net/http"

type Route struct {
	path       string
	controller func(responseWriter http.ResponseWriter, request *http.Request)
}

var ROUTES = []Route{
	{"GET /", index},
}
