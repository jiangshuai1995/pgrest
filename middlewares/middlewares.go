/*

@Time : 2019/9/14
@Author : Jiangs

*/
package middlewares

import (
	"github.com/dgrijalva/jwt-go"
	jwtmiddleware "github.com/iris-contrib/middleware/jwt"
	"github.com/kataras/iris"
	"github.com/kataras/iris/middleware/logger"
	"strings"
)

// JwtMiddleware check if actual request have JWT
func JwtMiddleware(key string, algo string) iris.Handler {
	jwtMiddleware := jwtmiddleware.New(jwtmiddleware.Config{
		ValidationKeyGetter: func(token *jwt.Token) (interface{}, error) {
			return []byte(key), nil
		},
		SigningMethod: jwt.GetSigningMethod(algo),
	})
	return jwtMiddleware.Serve
}

// Cors middleware
func Cors(origin []string, headers []string) iris.Handler{
	return func(ctx iris.Context) {
		ctx.Header("Access-Control-Allow-Origin",strings.Join(origin,","))
		ctx.Header("Access-Control-Allow-Credentials", "true")
		ctx.Header("Access-Control-Allow-Headers", "Access-Control-Allow-Origin,Content-Type")
		ctx.Header("Access-Control-Max-Age", "1800")
		ctx.Header("Access-Control-Allow-Methods", "GET,POST,PUT,DELETE,PATCH")
		ctx.Next()
	}
}

// Logger middleware
func Logs() (customLogger  iris.Handler){
	customLogger = logger.New(logger.Config{
		// Status displays status code
		Status: true,
		// IP displays request's remote address
		IP: true,
		// Method displays the http method
		Method: true,
		// Path displays the request path
		Path: true,
		// Query appends the url query to the Path.
		Query: true,

		//Columns: true,

		// if !empty then its contents derives from `ctx.Values().Get("logger_message")
		// will be added to the logs.
		//MessageContextKeys: []string{"logger_message"},

		// if !empty then its contents derives from `ctx.GetHeader("User-Agent")
		//MessageHeaderKeys: []string{"User-Agent"},
	})
	return
}
