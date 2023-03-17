package app

import (
	"fmt"

	"github.com/andrewarrow/feedback/router"
	"github.com/andrewarrow/feedback/sqlgen"
)

func workersByUserIdWhere(id int64) string {
	return fmt.Sprintf("user_id=%d", id)
}

func countWorkersByUserId(c *router.Context, id int64) int64 {
	where := workersByUserIdWhere(id)
	return c.Count("workers", where)
}

func addWorkerWithUserId(c *router.Context, id int64, jsonParams map[string]any) {
	jsonParams["user_id"] = id
	model := c.FindModel("worker")
	tableName := model.TableName()
	sql, params := sqlgen.InsertRow(tableName, model.Fields, jsonParams)
	c.Db.Exec(sql, params...)
}
