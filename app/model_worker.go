package app

import (
	"fmt"

	"github.com/andrewarrow/feedback/router"
)

func workersByUserIdWhere(id int64) string {
	return fmt.Sprintf("user_id=%d", id)
}

func workersByUserId(c *router.Context, id int64) []*map[string]any {
	where := workersByUserIdWhere(id)
	return c.SelectAllFrom("worker", "", where)
}

func countWorkersByUserId(c *router.Context, id int64) int64 {
	where := workersByUserIdWhere(id)
	return c.Count("worker", where)
}
