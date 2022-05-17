package middleware

import (
	"github.com/iris-contrib/middleware/jwt"
	"github.com/kataras/iris/v12"
)

var JWT *jwt.Middleware

func initJWT() {
	JWT = jwt.New(jwt.Config{
		ErrorHandler: func(context iris.Context, err error) {
			if err == nil {
				return
			}
			context.StopExecution()
			context.StatusCode(iris.StatusUnauthorized)
			//context.JSON(model.)
		},

		ValidationKeyGetter: func(token *jwt.Token) (interface{}, error) {
			return []byte("My Secret"), nil
		},

		SigningMethod: jwt.SigningMethodHS256,
	})
}
