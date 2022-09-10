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

	/*

		  {"data":
			  [{"name": "foo"},{"name": "bar"}]
			}

	*/
	fmt.Println(strings.Join(buff, "\n"))
}
