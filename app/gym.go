package app

import (
	"github.com/andrewarrow/feedback/router"
)

func HandleGym(c *router.Context, second, third string) {
	if second == "" && third == "" && c.Method == "GET" {
		handleGymIndex(c)
		return
	}
	c.NotFound = true
}

func handleGymIndex(c *router.Context) {

	send := map[string]any{}
	c.SendContentInLayout("welcome.html", send, 200)
}
