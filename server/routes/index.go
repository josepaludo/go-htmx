package server

import (
	"josepaludo/go-htmx/common/dir"
	"net/http"
)

func Index(responseWriter http.ResponseWriter, request *http.Request) {
	http.ServeFile(responseWriter, request, dir.Public("index.html"))
}

func Styles(responseWriter http.ResponseWriter, request *http.Request) {
	http.ServeFile(responseWriter, request, dir.Public("index.css"))
}

func Htmx(responseWriter http.ResponseWriter, request *http.Request) {
	http.ServeFile(responseWriter, request, dir.Public("htmx.min.js"))
}
