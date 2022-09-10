package gym

import (
	"api-gym/util"
	"fmt"
)

type Gym struct {
	Name   string
	Routes []*Route
}

func NewGym() *Gym {
	g := Gym{}
	g.Name = util.PseudoUuid()
	return &g
}

func (g *Gym) AddRoute(verb, route string) {
	g.Routes = append(g.Routes, NewRoute(verb, route))
}

func (g *Gym) ListRoutes() {
	for i, route := range g.Routes {
		fmt.Printf("%02d. %4s %s\n", i, route.Verb, route.Route)
	}
}
