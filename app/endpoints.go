package app

import (
	"strings"

	"codeberg.org/andrewarrow/roll"
	"github.com/andrewarrow/feedback/router"
	"github.com/andrewarrow/feedback/util"
)

var indexNameRoutes = "routes"
var workspaceGuid = "d841754e-1c41-44d8-91d1-1db0f3bb0f70"

func handleRouteUpsert(c *router.Context) {
	c.ReadJsonBodyIntoParams()
	c.Params["workspace"] = workspaceGuid
	id, _ := c.Params["id"].(string)
	guid := util.PseudoUuid()
	if id != "" {
		guid = id
	}
	roll.SingleUpsert(indexNameRoutes, guid, c.Params)

	route := c.Params["route"].(string)
	tokens := strings.Split(route, "/")
	top := tokens[1]

	send := map[string]any{}
	tops, m := getEndpoints(top)
	send["items"] = tops
	send["map"] = m
	c.SendContentAsJson(send, 200)
}
