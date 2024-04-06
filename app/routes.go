package app

import "github.com/andrewarrow/feedback/router"

func handleRoutes(c *router.Context) {
	send := map[string]any{}
	tops, m := getEndpoints()
	send["items"] = tops
	send["map"] = m
	c.SendContentAsJson(send, 200)
}
