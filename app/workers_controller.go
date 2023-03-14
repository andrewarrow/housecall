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
		c.SendContentInLayout("workers_index.html", nil, 200)
		return
	}
	handleWorkersCreate(c)
}

func handleWorkersCreate(c *router.Context) {
	c.NotFound = true
}

func handleWorkersShow(c *router.Context, id string) {
	if c.Method == "GET" {
		c.SendContentInLayout("workers_show.html", nil, 200)
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