package simulate

import (
	"api-gym/gym"
)

func Json(model string, g *gym.Gym) {
	s := g.StructsByName[model]
	PrintItemsToStdout(s, g)
}
