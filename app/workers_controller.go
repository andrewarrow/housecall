package app

import (
	"github.com/andrewarrow/feedback/router"
	"github.com/andrewarrow/feedback/sqlgen"
)

func HandleWorkers(c *router.Context, second, third string) {
	if c.User == nil {
		c.UserRequired = true
		return
	}
	if second == "" {
		handleWorkersIndex(c)
	} else if second != "" && third == "" {
		handleWorkersShow(c, second)
	} else {
		c.NotFound = true
	}
}

func handleWorkersIndex(c *router.Context) {
	if c.Method == "GET" {
		rows := c.SelectAllFrom("workers", "", "")
		c.SendContentInLayout("workers_index.html", rows, 200)
		return
	}
	handleWorkersCreate(c)
}

func handleWorkersCreate(c *router.Context) {
	//params := c.ReadBodyIntoJson()

	model := c.FindModel("worker")
	tableName := model.TableName()
	sql, params := sqlgen.InsertRow(tableName, model.Fields)
	c.Db.Exec(sql, params...)

	rows := c.SelectAllFrom("workers", "", "")
	c.ExecuteTemplate("workers_list.html", rows)
}

func handleWorkersShow(c *router.Context, id string) {
	if c.Method == "GET" {
		row := c.SelectOneFrom(id, "workers")
		c.SendContentInLayout("workers_show.html", row, 200)
		return
	}
	handleWorkersUpdates(c, id)
}

func handleWorkersUpdates(c *router.Context, id string) {
	if c.Method == "POST" {
		c.NotFound = true
	} else if c.Method == "DELETE" {
		c.NotFound = true
	}
}
