package server

import (
	"api-gym/files"
	"api-gym/gym"
	"api-gym/simulate"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Setup(g *gym.Gym) *gin.Engine {

	router := gin.Default()
	for _, route := range g.Routes {
		s := g.StructsByName[route.Model]
		j := JsonRoute{s, g, route}
		if route.Verb == "GET" {
			router.GET(route.Route, j.JsonRunner)
		} else if route.Verb == "POST" {
			router.POST(route.Route, j.JsonRunner)
		}
	}
	return router
}

type JsonRoute struct {
	Struct *gym.Struct
	Gym    *gym.Gym
	Route  *gym.Route
}

func (j *JsonRoute) JsonRunner(c *gin.Context) {
	c.Writer.Header().Set("Content-Type", "application/json")
	stringToSend := ""
	if j.Route.UseFile != "" {
		stringToSend = files.ReadFile("static/" + j.Route.UseFile)
	} else {
		stringToSend = simulate.PrintItemsToString(j.Struct, j.Gym, j.Route.Count)
	}
	c.String(http.StatusOK, stringToSend)
}
