package router

import (
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/context"
)

func Route(app *iris.Application) {
	app.Get("/", func(context *context.Context) {
		context.HTML("<h1>" + "FITNESS" + "</h1>")
	})
}
