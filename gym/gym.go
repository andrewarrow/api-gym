package gym

import (
	"api-gym/files"
	"api-gym/util"
	"encoding/json"
	"fmt"
)

type Gym struct {
	Name   string   `json:"name"`
	Routes []*Route `json:"routes"`
}

func NewGym() *Gym {
	g := Gym{}
	g.Name = util.PseudoUuid()
	return &g
}

func (g *Gym) AddRoute(verb, route string) {
	g.Routes = append(g.Routes, NewRoute(verb, route))
	g.Save()
}

func (g *Gym) ListRoutes() {
	for i, route := range g.Routes {
		fmt.Printf("%02d. %4s %s\n", i, route.Verb, route.Route)
	}
}

func LoadGym() *Gym {
	gymJson := files.ReadFile("gym.json")
	if gymJson == "" {
		return NewGym()
	}
	var g Gym
	json.Unmarshal([]byte(gymJson), &g)
	return &g
}

func (g *Gym) Save() {
	b, _ := json.Marshal(g)
	files.SaveFile("gym.json", string(b))
}
