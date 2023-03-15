package app

import (
	"github.com/andrewarrow/feedback/router"
)

func HandleCustomers(c *router.Context, second, third string) {
	if c.User == nil {
		c.UserRequired = true
		return
	}
	if second == "" {
		handleCustomersIndex(c)
	} else if second != "" && third == "" {
		handleCustomersShow(c, second)
	} else {
		c.NotFound = true
	}
}

func handleCustomersIndex(c *router.Context) {
	if c.Method == "GET" {
	  rows := c.SelectAllFrom("customers")
		c.SendContentInLayout("customers_index.html", rows, 200)
		return
	}
	handleCustomersCreate(c)
}

func handleCustomersCreate(c *router.Context) {
	c.NotFound = true
}

func handleCustomersShow(c *router.Context, id string) {
	if c.Method == "GET" {
		row := c.SelectOneFrom(id, "customers")
		c.SendContentInLayout("customers_show.html", row, 200)
		return
	}
	handleCustomersUpdates(c, id)
}

func handleCustomersUpdates(c *router.Context, id string) {
	if c.Method == "POST" {
		c.NotFound = true
	} else if c.Method == "DELETE" {
		c.NotFound = true
	}
}