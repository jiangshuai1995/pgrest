/*

@Time : 2019/9/14
@Author : Jiangs

*/
package middlewares

import (
	"github.com/kataras/iris"
)
var (
	app *iris.Application

	// MiddlewareStack on pgREST
	MiddlewareStack []iris.Handler
)

func initApp(){
	if len(MiddlewareStack) == 0 {
		MiddlewareStack = append(MiddlewareStack)
		MiddlewareStack = append(MiddlewareStack,Logs())
		//if !config.PrestConf.Debug && config.PrestConf.EnableDefaultJWT {
		//	MiddlewareStack = append(MiddlewareStack, JwtMiddleware(config.PrestConf.JWTKey, config.PrestConf.JWTAlgo))
		//}
		//if config.PrestConf.CORSAllowOrigin != nil {
		//	MiddlewareStack = append(MiddlewareStack, Cors(config.PrestConf.CORSAllowOrigin, config.PrestConf.CORSAllowHeaders))
		//}
		//MiddlewareStack = append(MiddlewareStack,JwtMiddleware("secret","HS256"))

		//MiddlewareStack = append(MiddlewareStack,Cors(a,b))
	}
	//a:=[]string{"http://foo.com"}
	//b:=[]string{"*"}
	app = iris.New()
	app.Use(MiddlewareStack...)
}

func GetApp() *iris.Application{
	if app == nil {
		initApp()
	}
	return app
}

