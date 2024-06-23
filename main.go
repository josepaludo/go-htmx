package main

import (
	"josepaludo/go-htmx/src/dir"
	"josepaludo/go-htmx/src/env"
	"josepaludo/go-htmx/src/server"
)

func main() {

	env.Init()
	dir.Init()

	server.Serve()
}
