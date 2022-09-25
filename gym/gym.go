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
	g.AddStruct("user")
	g.AddFieldToStruct("user", "id", "uuid", "")
	g.AddFieldToStruct("user", "first_name", "first_name", "")
	g.AddFieldToStruct("user", "last_name", "last_name", "")
	g.AddFieldToStruct("user", "pronouns", "pronouns", "")
	g.AddFieldToStruct("user", "email", "email", "")
	g.AddFieldToStruct("user", "phone", "phone", "")
	g.AddFieldToStruct("user", "age", "int", "max:100")
	return &g
}

func (g *Gym) AddFieldToStruct(model, name, flavor, extra string) {
	var field *Field
	s := g.StructsByName[model]
	field = NewField(name, flavor, extra)
	s.Fields = append(s.Fields, field)
	g.Save()
}

func (g *Gym) AddStruct(name string) {
	g.Structs = append(g.Structs, NewStruct(name))
	g.fillStructsByName()
	pluralName := name + "s" // TODO handle company companies case
	g.AddRoute("GET", fmt.Sprintf("/api/v1/%s", pluralName), name, 10)
	g.AddRoute("GET", fmt.Sprintf("/api/v1/%s/:id", pluralName), name, 1)
	g.Save()
}

func (g *Gym) RemoveStruct(name string) {
	newList := []*Struct{}
	for _, s := range g.Structs {
		if s.Name == name {
			continue
		}
		newList = append(newList, s)
	}
	g.Structs = newList
	delete(g.StructsByName, name)
	g.Save()
}

func (g *Gym) AddRoute(verb, route, model string, count int) {
	g.Routes = append(g.Routes, NewRoute(verb, route, model, count))
	g.SortRoutes()
	g.Save()
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
	/*
		for i, route := range g.Routes {
			fmt.Printf("%2d. %-30s %s\n", i+1, route.String(), route.Response)
		}
		fmt.Println("")
		for i, s := range g.Structs {
			fmt.Printf("%2d. %-30s\n", i+1, s.Name)
			for j, f := range s.Fields {
				fmt.Printf("  %2d. %-20s %-26s %-16s %s\n", j+1, f.Name, f.Flavor, f.Random, f.Extra)
			}
		}*/
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
	//indexAsInt, _ := strconv.Atoi(index)
	//g.Routes[indexAsInt].Response = response
	//g.Structs = append(g.Structs, NewStruct(response))
	//g.Save()
	//g.ListRoutes()
}

func (g *Gym) DeleteField(modelIndex, fieldIndex int) {
	newList := []*Field{}
	for i, f := range g.Structs[modelIndex].Fields {
		if i == fieldIndex {
			continue
		}
		newList = append(newList, f)
	}
	g.Structs[modelIndex].Fields = newList
	g.Save()
}

func (g *Gym) UpdateStructRandom(index, fieldIndex, random string) {
	/*
		indexAsInt, _ := strconv.Atoi(index)
		fieldIndexAsInt, _ := strconv.Atoi(fieldIndex)
		g.Structs[indexAsInt].Fields[fieldIndexAsInt].Random = random
		g.Save()
		g.ListRoutes()
	*/
}
