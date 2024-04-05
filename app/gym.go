package app

import (
	"github.com/andrewarrow/feedback/router"
)

func HandleGym(c *router.Context, second, third string) {
	if second == "" && third == "" && c.Method == "GET" {
		handleGymIndex(c)
		return
	}
	if second == "json" && third == "" && c.Method == "GET" {
		handleGymJson(c)
		return
	}
	c.NotFound = true
}

func handleGymIndex(c *router.Context) {

	send := map[string]any{}
	send["items"] = []string{"", ""}
	c.SendContentInLayout("endpoints.html", send, 200)
}

func handleGymJson(c *router.Context) {

	send := map[string]any{}
	c.SendContentInLayout("gym.html", send, 200)
}
