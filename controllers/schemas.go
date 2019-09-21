package controllers

import (
	"fmt"
	"github.com/kataras/iris"
	"strings"

	"pgrest/config"
)

// GetSchemas list all (or filter) schemas
func GetSchemas(ctx iris.Context) {
	r:= ctx.Request()
	requestWhere, values, err := config.PrestConf.Adapter.WhereByRequest(r, 1)
	if err != nil {
		ctx.StatusCode(iris.StatusBadRequest)
		ctx.Recorder().WriteString(err.Error())
		ctx.Next()
		return
	}

	sqlSchemas, hasCount := config.PrestConf.Adapter.SchemaClause(r)

	if requestWhere != "" {
		sqlSchemas = fmt.Sprint(sqlSchemas, " WHERE ", requestWhere)
	}

	distinct, err := config.PrestConf.Adapter.DistinctClause(r)
	if err != nil {
		ctx.StatusCode(iris.StatusBadRequest)
		ctx.Recorder().WriteString(err.Error())
		ctx.Next()
		return
	}
	if distinct != "" {
		sqlSchemas = strings.Replace(sqlSchemas, "SELECT", distinct, 1)
	}

	order, err := config.PrestConf.Adapter.OrderByRequest(r)
	if err != nil {
		ctx.StatusCode(iris.StatusBadRequest)
		ctx.Recorder().WriteString(err.Error())
		ctx.Next()
		return
	}
	order = config.PrestConf.Adapter.SchemaOrderBy(order, hasCount)

	page, err := config.PrestConf.Adapter.PaginateIfPossible(r)
	if err != nil {
		ctx.StatusCode(iris.StatusBadRequest)
		ctx.Recorder().WriteString(err.Error())
		ctx.Next()
		return
	}

	sqlSchemas = fmt.Sprint(sqlSchemas, order, " ", page)
	sc := config.PrestConf.Adapter.Query(sqlSchemas, values...)
	if sc.Err() != nil {
		ctx.StatusCode(iris.StatusBadRequest)
		ctx.Recorder().WriteString(err.Error())
		ctx.Next()
		return
	}
	ctx.Header("Content-Type", "application/json")
	ctx.Write(sc.Bytes())
}
