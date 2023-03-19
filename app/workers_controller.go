package app

import (
	"github.com/andrewarrow/feedback/router"
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
		rows := c.SelectAll("worker", "", []any{})
		c.SendContentInLayout("workers_index.html", rows, 200)
		return
	}
	handleWorkersCreate(c)
}

func handleWorkersCreate(c *router.Context) {
	jsonParams := c.ReadBodyIntoJson()

	addWorkerWithUserId(c, c.User["id"].(int64), jsonParams)

	params := []any{c.User["id"]}
	rows := c.SelectAll("worker", "user_id=$1", params)
	c.ExecuteTemplate("workers_list.html", rows)
}

func handleWorkersShow(c *router.Context, id string) {
	if c.Method == "GET" {
		params := []any{id}
		row := c.SelectOne("worker", "where guid=$1", params)
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
