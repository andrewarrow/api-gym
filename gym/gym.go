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
	Name          string             `json:"name"`
	Routes        []*Route           `json:"routes"`
	Structs       []*Struct          `json:"structs"`
	StructsByName map[string]*Struct `json:"-"`
}

func NewGym() *Gym {
	g := Gym{}
	g.Name = util.PseudoUuid()
	return &g
}

func (g *Gym) AddRoute(verb, route, modelIndex string) {
	g.Routes = append(g.Routes, NewRoute(verb, route, modelIndex))
	g.SortRoutes()
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
	g.SortRoutes()
	g.Save()
	g.ListRoutes()
}

func (g *Gym) SortRoutes() {
	sort.SliceStable(g.Routes, func(i, j int) bool {
		return g.Routes[i].String() < g.Routes[j].String()
	})
}

func (g *Gym) ListRoutes() {
	fmt.Println("")
	for i, route := range g.Routes {
		fmt.Printf("%2d. %-30s %s\n", i+1, route.String(), route.Response)
	}
	fmt.Println("")
	for i, s := range g.Structs {
		fmt.Printf("%2d. %-30s\n", i+1, s.Name)
		for j, f := range s.Fields {
			fmt.Printf("  %2d. %-20s %-26s %-16s %s\n", j+1, f.Name, f.Flavor, f.Random, f.Extra)
		}
	}
	fmt.Println("")
}

func LoadGym() *Gym {
	gymJson := files.ReadFile("gym.json")
	if gymJson == "" {
		return NewGym()
	}
	var g Gym
	json.Unmarshal([]byte(gymJson), &g)
	g.fillStructsByName()
	return &g
}

func (g *Gym) fillStructsByName() {
	g.StructsByName = map[string]*Struct{}
	for _, s := range g.Structs {
		g.StructsByName[s.Name] = s
	}
}

func (g *Gym) Save() {
	b, _ := json.Marshal(g)
	files.SaveFile("gym.json", string(b))
}
func (g *Gym) SaveBackup() {
	b, _ := json.Marshal(g)
	files.SaveFile("gym.backup.json", string(b))
}

func (g *Gym) AddResponseToRoute(index, response string) {
	indexAsInt, _ := strconv.Atoi(index)
	g.Routes[indexAsInt].Response = response
	g.Structs = append(g.Structs, NewStruct(response))
	g.Save()
	g.ListRoutes()
}

func (g *Gym) AddFieldToStruct(modelIndex, flavorIndex int) {
	var field *Field
	field = NewField("name", "string", "timestamp", "")
	g.Structs[modelIndex].Fields = append(g.Structs[modelIndex].Fields, field)
	g.Save()
}

func (g *Gym) AddStruct(name string) {
	g.Structs = append(g.Structs, NewStruct(name))
	g.fillStructsByName()
	g.Save()
	g.ListRoutes()
}

func (g *Gym) UpdateStructRandom(index, fieldIndex, random string) {
	indexAsInt, _ := strconv.Atoi(index)
	fieldIndexAsInt, _ := strconv.Atoi(fieldIndex)
	g.Structs[indexAsInt].Fields[fieldIndexAsInt].Random = random
	g.Save()
	g.ListRoutes()
}
