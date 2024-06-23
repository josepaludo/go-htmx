package server

import (
	"josepaludo/go-htmx/common/env"
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
