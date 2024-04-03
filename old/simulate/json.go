package simulate

import (
	"api-gym/gym"
)

func Json(model string, g *gym.Gym) {
	s := g.StructsByName[model]
	/*
		for k, v := range g.StructsByName {
			fmt.Println(k, v)
		}*/
	PrintItemsToStdout(s, g)
}
