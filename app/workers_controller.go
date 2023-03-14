package app

import (
  "github.com/andrewarrow/feedback/router"
)

func HandleWorkers(c *router.Context, second, third string) {
  if c.User == nil {
    c.UserRequired = true
    return
  }
  c.SendContentInLayout("workers_index.html", nil, 200)
}