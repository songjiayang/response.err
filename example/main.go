package main

import (
	"github.com/dolab/gogo"
	"github.com/songjiayang/response.err"
)

func main() {
	app := gogo.New("development", "")

	// GET /
	app.GET("/", func(ctx *gogo.Context) {
		ctx.Text("Hello, gogo!")
	})

	app.GET("/err", func(ctx *gogo.Context) {
		ctx.Json(errors.NewErrorResponse(ctx.RequestID(), ctx.RequestURI(), errors.InvalidParameter))
	})

	app.Run()
}
