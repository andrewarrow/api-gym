package simulate

import (
	"api-gym/gym"
	"strconv"
)

func Json(index, extra string, g *gym.Gym) {
	indexAsInt, _ := strconv.Atoi(index)
	s := g.Structs[indexAsInt-1]
	PrintItemsToStdout(s, extra, g)
}
