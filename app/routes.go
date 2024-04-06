package app

import "github.com/andrewarrow/feedback/router"

func handleRoutes(c *router.Context) {
	top := c.Request.URL.Query().Get("top")
	send := map[string]any{}
	tops, m := getEndpoints(top)
	send["items"] = tops
	send["map"] = m
	c.SendContentAsJson(send, 200)
}
