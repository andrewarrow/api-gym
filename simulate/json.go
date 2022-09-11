package simulate

import (
	"api-gym/gym"
	"strconv"
)

func Json(index, amount string, g *gym.Gym) {
	indexAsInt, _ := strconv.Atoi(index)
	amountAsInt, _ := strconv.Atoi(amount)
	s := g.Structs[indexAsInt-1]
	printItems(s, amountAsInt, g)
}
