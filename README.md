# pgrest

Fork from [prest](https://github.com/prest/prest)

## Postgres version

- 9.4 or higher

## Modify

- 为了使用prest的扩展功能，对源码进行修改后提交至本仓库。
- 删除了SQL Scripts(_QUERIES)，Migrations，Version

## Docs

https://postgres.rest/ ([source](https://github.com/prest/prest.github.io))


## Tolist

- 未使用分组路由,使用app.Done(),但是前面的handler都需要使用ctx.Next()
- 错误处理未完成所有代码
    >现未使用OnErrorCode()方法，因为OnErrorCode()前会resetRcorder
- Content-Type需在write前写入,格式转换方法需完善
- mux方法全部转换未iris方法

## Done

- 对指定数据库和模式下表数量的查询
