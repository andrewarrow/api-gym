package app

import (
	"strings"

	"github.com/andrewarrow/feedback/router"
)

var indexNameRoutes = "routes"
var workspaceGuid = "d841754e-1c41-44d8-91d1-1db0f3bb0f70"

func handleRouteUpsert(c *router.Context) {
	c.ReadJsonBodyIntoParams()
	c.Params["workspace"] = workspaceGuid
	c.Params["name"] = c.Params["route"]
	id, _ := c.Params["id"].(string)
	if id == "" {
		c.ValidateCreate("route")
		c.Insert("route")
	} else {
		c.Update("route", "where guid=", id)
	}

	route := c.Params["route"].(string)
	tokens := strings.Split(route, "/")
	top := tokens[1]

	send := map[string]any{}
	tops, m := getEndpoints(c, top)
	send["items"] = tops
	send["map"] = m
	c.SendContentAsJson(send, 200)
}
