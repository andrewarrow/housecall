package app

import (
	"github.com/andrewarrow/feedback/router"
)

func HandleWelcome(c *router.Context, second, third string) {
	c.SendContentInLayout("welcome.html", nil, 200)
}
