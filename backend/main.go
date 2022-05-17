package main

import (
	"fitness/config"
	"fitness/router"
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/middleware/logger"
	"github.com/kataras/iris/v12/middleware/recover"
)

func main() {
	app := iris.New()

	// Set logger level
	app.Logger().SetLevel(config.Viper.GetString("server.logger.level"))

	// Add recover to recover from any http-relative panics
	app.Use(recover.New())

	// Add logger to log the requests to the terminal
	app.Use(logger.New())

	// Globally allow options to enable CORS
	app.AllowMethods(iris.MethodOptions)

	// Add global CORS handler
	//app.Use(middleware.CORS)

	// Router
	router.Route(app)

	_ = app.Run(iris.Addr(config.Viper.GetString("server.addr")), iris.WithoutServerError(iris.ErrServerClosed))
}
