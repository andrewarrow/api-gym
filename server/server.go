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
	/*
		for _, route := range g.Routes {
			modelIndexAsInt, _ := strconv.Atoi(route.ModelIndex)
			s := g.Structs[modelIndexAsInt-1]
			j := JsonRoute{s, g, route.UseFile}
			if route.Verb == "get" {
				router.GET(route.Route, j.JsonRunner)
			} else if route.Verb == "post" {
				router.POST(route.Route, j.JsonRunner)
			}
		}
	*/
	return router
}

type JsonRoute struct {
	Struct  *gym.Struct
	Gym     *gym.Gym
	UseFile string
}

func (j *JsonRoute) JsonRunner(c *gin.Context) {
	c.Writer.Header().Set("Content-Type", "application/json")
	stringToSend := ""
	if j.UseFile != "" {
		stringToSend = files.ReadFile("static/" + j.UseFile)
	} else {
		stringToSend = simulate.PrintItemsToString(j.Struct, j.Gym)
	}
	c.String(http.StatusOK, stringToSend)
}
