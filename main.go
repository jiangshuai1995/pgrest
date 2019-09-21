/*

@Time : 2019/9/14
@Author : Jiangs

*/
package main

import (
	"github.com/kataras/iris"
	"pgrest/adapters/postgres"
	"pgrest/config"
	"pgrest/controllers"
	"pgrest/middlewares"
)

func main() {
	config.Load()

	postgres.Load()

	app:=middlewares.GetApp()
	app.Done(middlewares.FormatMiddleware())
	app.Get("/version", Versionhandler)
	app.Get("/databases",controllers.GetDatabases)
	app.Get("/schemas",controllers.GetSchemas)
	app.Get("/tables",controllers.GetTables)
	app.Get("/{database}/{schema}",controllers.GetTablesByDatabaseAndSchema)
	app.Get("/{database}/{schema}/{table}",controllers.SelectFromTables)

	app.Run(iris.Addr(":9000"))

}

func Versionhandler(ctx iris.Context) {
	ctx.WriteString("Hello World")

}
