package simulate

import (
	"api-gym/gym"
	"fmt"
	"strconv"
)

func Run(index string, g *gym.Gym) {
	indexAsInt, _ := strconv.Atoi(index)
	route := g.Routes[indexAsInt]
	fmt.Println("Running...", route)
}
