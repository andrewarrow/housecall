package app

import (
	"github.com/andrewarrow/feedback/router"
)

func HandleAppointments(c *router.Context, second, third string) {
	if c.User == nil {
		c.UserRequired = true
		return
	}
	if second == "" {
		handleAppointmentsIndex(c)
	} else if second != "" && third == "" {
		handleAppointmentsShow(c, second)
	} else {
		c.NotFound = true
	}
}

func handleAppointmentsIndex(c *router.Context) {
	if c.Method == "GET" {
		params := []any{}
		rows := c.SelectAll("appointment", "where appointment_at > NOW() order by appointment_at asc", params)
		c.SendContentInLayout("appointments_index.html", rows, 200)
		return
	}
	handleAppointmentsCreate(c)
}

func handleAppointmentsCreate(c *router.Context) {
	c.NotFound = true
}

func handleAppointmentsShow(c *router.Context, id string) {
	if c.Method == "GET" {
		params := []any{id}
		row := c.SelectOne("appointment", "guid=$1", params)
		c.SendContentInLayout("appointments_show.html", row, 200)
		return
	}
	handleAppointmentsUpdates(c, id)
}

func handleAppointmentsUpdates(c *router.Context, id string) {
	if c.Method == "POST" {
		c.NotFound = true
	} else if c.Method == "DELETE" {
		c.NotFound = true
	}
}
