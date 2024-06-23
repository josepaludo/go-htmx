package server

import (
	"josepaludo/go-htmx/src/dir"
	"josepaludo/go-htmx/src/env"
	"net/http"
)

func Serve() {

	handler := http.NewServeMux()

	for _, route := range ROUTES {
		handler.HandleFunc(route.path, route.controller)
	}

	server := http.Server{
		Addr:    env.API_PORT,
		Handler: handler,
	}

	if err := server.ListenAndServe(); err != nil {
		panic(err)
	}
}

func index(responseWriter http.ResponseWriter, request *http.Request) {
	http.ServeFile(responseWriter, request, dir.Static("index.html"))
}
