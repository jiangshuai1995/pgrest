package controllers

import (
	"fmt"
	"github.com/kataras/iris"
	"strings"

	"pgrest/config"
)

// GetDatabases list all (or filter) databases
func GetDatabases(ctx iris.Context) {
	r := ctx.Request()
	requestWhere, values, err := config.PrestConf.Adapter.WhereByRequest(r, 1)
	if err != nil {
		ctx.StatusCode(iris.StatusBadRequest)
		ctx.Recorder().WriteString(err.Error())
		ctx.Next()
		return
	}
	requestWhere = config.PrestConf.Adapter.DatabaseWhere(requestWhere)

	query, hasCount := config.PrestConf.Adapter.DatabaseClause(r)
	sqlDatabases := fmt.Sprint(query, requestWhere)
	distinct, err := config.PrestConf.Adapter.DistinctClause(r)
	if err != nil {
		ctx.StatusCode(iris.StatusBadRequest)
		ctx.Recorder().WriteString(err.Error())
		ctx.Next()
		return
	}
	if distinct != "" {
		sqlDatabases = strings.Replace(sqlDatabases, "SELECT", distinct, 1)
	}

	order, err := config.PrestConf.Adapter.OrderByRequest(r)
	if err != nil {
		ctx.StatusCode(iris.StatusBadRequest)
		ctx.Recorder().WriteString(err.Error())
		ctx.Next()
		return
	}
	order = config.PrestConf.Adapter.DatabaseOrderBy(order, hasCount)

	sqlDatabases = fmt.Sprint(sqlDatabases, order)

	page, err := config.PrestConf.Adapter.PaginateIfPossible(r)
	if err != nil {
		ctx.StatusCode(iris.StatusBadRequest)
		ctx.Recorder().WriteString(err.Error())
		ctx.Next()
		return
	}

	sqlDatabases = fmt.Sprint(sqlDatabases, " ", page)
	sc := config.PrestConf.Adapter.Query(sqlDatabases, values...)
	if sc.Err() != nil {
		ctx.StatusCode(iris.StatusBadRequest)
		ctx.Recorder().WriteString(err.Error())
		ctx.Next()
		return
	}
	ctx.Header("Content-Type", "application/json")
	ctx.Write(sc.Bytes())
}
