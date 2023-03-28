package app

import (
	"fmt"

	"github.com/andrewarrow/feedback/router"
)

func workersByUserIdWhere(id int64) string {
	return fmt.Sprintf("user_id=%d", id)
}

func countWorkersByUserId(c *router.Context, id int64) int64 {
	where := workersByUserIdWhere(id)
	return c.Count("workers", where)
}

func addWorkerWithUserId(c *router.Context, id int64, jsonParams map[string]any) {
	c.Params = jsonParams
	c.Validate("worker")
	c.Insert("worker")
}
