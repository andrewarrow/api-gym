package app

import "github.com/andrewarrow/feedback/router"

var indexNameRoutes = "routes"

func handleRouteUpsert(c *router.Context) {

	c.SendContentAsJsonMessage("ok", 200)
}
