package main

import (
	"josepaludo/go-htmx/common/dir"
	"josepaludo/go-htmx/common/env"
	"josepaludo/go-htmx/server"
)

func main() {

	env.Init()
	dir.Init()

	server.Serve()
}
