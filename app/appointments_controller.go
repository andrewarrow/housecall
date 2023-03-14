package app

import (
  "github.com/andrewarrow/feedback/router"
)

func HandleAppointments(c *router.Context, second, third string) {
  if c.User == nil {
    c.UserRequired = true
    return
  }
  c.SendContentInLayout("appointments_index.html", nil, 200)
}