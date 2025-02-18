package main

import (
	"gofr-test/http/handler"

	"gofr.dev/pkg/gofr"
)

func main() {
	app := gofr.New()

	h := handler.NewHttpDi()
	uh := h.Handler.UserHandler
	app.GET("/user", uh.GetUserHandler)

	app.Run()
}
