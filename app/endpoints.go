package app

import (
	"codeberg.org/andrewarrow/roll"
	"github.com/andrewarrow/feedback/router"
	"github.com/andrewarrow/feedback/util"
)

var indexNameRoutes = "routes"
var workspaceGuid = "d841754e-1c41-44d8-91d1-1db0f3bb0f70"

func handleRouteUpsert(c *router.Context) {
	c.ReadJsonBodyIntoParams()
	c.Params["workspace"] = workspaceGuid
	guid := util.PseudoUuid()
	roll.SingleUpsert(indexNameRoutes, guid, c.Params)

	c.SendContentAsJsonMessage("ok", 200)
}
