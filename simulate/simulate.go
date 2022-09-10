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

func printItems(flavor string, g *gym.Gym) {
	buff := []string{}
	buff = append(buff, `{"data":`)

	s := g.StructsByName[flavor]
	buff = append(buff, "[")
	sub := []string{}
	for i := 0; i < 9; i++ {
		thing := []string{}
		items := []string{}
		thing = append(thing, "{")
		for _, f := range s.Fields {
			items = append(items, fmt.Sprintf(`"%s": "%s"`, f.Name, "foo"))
		}
		thing = append(thing, strings.Join(items, ","))
		thing = append(thing, "}")
		sub = append(sub, strings.Join(thing, ""))
	}
	buff = append(buff, strings.Join(sub, ","))
	buff = append(buff, "]")
	buff = append(buff, "}")

	/*

		  {"data":
			  [{"id": "123", "name": "foo"},{"name": "bar"}]
			}

	*/
	fmt.Println(strings.Join(buff, "\n"))
}
