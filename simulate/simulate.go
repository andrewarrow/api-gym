package simulate

import (
	"api-gym/gym"
	"fmt"
	"strconv"
	"strings"
)

func Run(index string, g *gym.Gym) {
	indexAsInt, _ := strconv.Atoi(index)
	route := g.Routes[indexAsInt]
	structure, flavor := route.ParseResponse()

	if structure == "array" {
		printItems(flavor, g)
		return
	}
}

func makeArrayItems(field *gym.Field, g *gym.Gym) string {
	sub := []string{}
	//for i := 0; i < 6; i++ {
	//sub = append(sub, foo(s))
	//	}
	return strings.Join(sub, "")
}

func printItems(flavor string, g *gym.Gym) {
	buff := []string{}
	buff = append(buff, `{"data":`)

	s := g.StructsByName[flavor]
	buff = append(buff, "[")
	sub := []string{}
	for i := 0; i < 9; i++ {
		sub = append(sub, makeStructJson(s))
	}
	buff = append(buff, strings.Join(sub, ","))
	buff = append(buff, "]")
	buff = append(buff, "}")

	fmt.Println(strings.Join(buff, "\n"))
}

func makeStructJson(s *gym.Struct) string {
	thing := []string{}
	items := []string{}
	thing = append(thing, "{")
	for _, f := range s.Fields {
		if f.Flavor == "string" {
			items = append(items, fmt.Sprintf(`"%s": "%s"`, f.NameToJson(), f.ToFakeValue()))
		} else if strings.HasPrefix(f.Flavor, "[]") {
			//items = append(items, makeArrayItems(f))
		} else {
			items = append(items, fmt.Sprintf(`"%s": %s`, f.NameToJson(), f.ToFakeValue()))
		}
	}
	thing = append(thing, strings.Join(items, ","))
	thing = append(thing, "}")
	return strings.Join(thing, "")
}
