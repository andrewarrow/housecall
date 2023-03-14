package app

import (
	"github.com/andrewarrow/feedback/router"
)

func HandleDashboard(c *router.Context, second, third string) {
	if c.User == nil {
		c.UserRequired = true
		return
	}
	vars := map[string]int{"workers": 2, "appointments": 7}
	c.SendContentInLayout("dashboard.html", vars, 200)
}
