package server

import (
	"api-gym/gym"
	"api-gym/simulate"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func Setup(g *gym.Gym) *gin.Engine {

	router := gin.Default()

	for _, route := range g.Routes {
		modelIndexAsInt, _ := strconv.Atoi(route.ModelIndex)
		s := g.Structs[modelIndexAsInt-1]
		j := JsonRoute{s, g}
		if route.Verb == "get" {
			router.GET(route.Route, j.JsonRunner)
		} else if route.Verb == "post" {
			router.POST(route.Route, j.JsonRunner)
		}
	}
	return router
}

type JsonRoute struct {
	Struct *gym.Struct
	Gym    *gym.Gym
}

func (j *JsonRoute) JsonRunner(c *gin.Context) {
	jsonString := simulate.PrintItemsToString(j.Struct, j.Struct.Extra, j.Gym)
	var m map[string]interface{}
	json.Unmarshal([]byte(jsonString), &m)
	c.JSON(http.StatusOK, m)
}
