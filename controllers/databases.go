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
		ctx.WriteString(err.Error())
		//http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	requestWhere = config.PrestConf.Adapter.DatabaseWhere(requestWhere)

	query, hasCount := config.PrestConf.Adapter.DatabaseClause(r)
	sqlDatabases := fmt.Sprint(query, requestWhere)
	distinct, err := config.PrestConf.Adapter.DistinctClause(r)
	if err != nil {
		ctx.StatusCode(iris.StatusBadRequest)
		ctx.WriteString(err.Error())
		//http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	if distinct != "" {
		sqlDatabases = strings.Replace(sqlDatabases, "SELECT", distinct, 1)
	}

	order, err := config.PrestConf.Adapter.OrderByRequest(r)
	if err != nil {
		ctx.StatusCode(iris.StatusBadRequest)
		ctx.WriteString(err.Error())
		//http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	order = config.PrestConf.Adapter.DatabaseOrderBy(order, hasCount)

	sqlDatabases = fmt.Sprint(sqlDatabases, order)

	page, err := config.PrestConf.Adapter.PaginateIfPossible(r)
	if err != nil {
		ctx.StatusCode(iris.StatusBadRequest)
		ctx.WriteString(err.Error())
		//http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	sqlDatabases = fmt.Sprint(sqlDatabases, " ", page)
	sc := config.PrestConf.Adapter.Query(sqlDatabases, values...)
	if sc.Err() != nil {
		//http.Error(w, sc.Err().Error(), http.StatusBadRequest)
		return
	}
	ctx.Write(sc.Bytes())
}
