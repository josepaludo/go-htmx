package server

import (
	"josepaludo/go-htmx/common/dir"
	"net/http"
)

func Index(responseWriter http.ResponseWriter, request *http.Request) {
	http.ServeFile(responseWriter, request, dir.Static("index.html"))
}
