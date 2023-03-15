package app

import (
	"github.com/andrewarrow/feedback/router"
)

func HandleDashboard(c *router.Context, second, third string) {
	if c.User == nil {
		c.UserRequired = true
		return
	}
	workersCount := c.Count("workers")
	appointmentsCount := c.Count("appointments")

	vars := map[string]any{"workers": workersCount, "appointments": appointmentsCount}
	c.SendContentInLayout("dashboard.html", vars, 200)
}
