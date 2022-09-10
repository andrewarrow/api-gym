package gym

import (
	"api-gym/files"
	"api-gym/util"
	"encoding/json"
	"fmt"
	"sort"
	"strconv"
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
	g.ListRoutes()
}

func (g *Gym) RemoveRoute(index string) {
	indexAsInt, _ := strconv.Atoi(index)
	routes := []*Route{}
	for i, route := range g.Routes {
		if i == indexAsInt {
			continue
		}
		routes = append(routes, route)
	}
	g.Routes = routes
	g.Save()
	g.ListRoutes()
}

func (g *Gym) ListRoutes() {
	buff := []string{}
	for _, route := range g.Routes {
		buff = append(buff, fmt.Sprintf("%4s %s", route.Verb, route.Route))
	}
	sort.SliceStable(buff, func(i, j int) bool {
		return buff[i] < buff[j]
	})
	for i, route := range buff {
		fmt.Printf("%2d. %s\n", i, route)
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
func (g *Gym) SaveBackup() {
	b, _ := json.Marshal(g)
	files.SaveFile("gym.backup.json", string(b))
}
