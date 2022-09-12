package flavor

import (
	"fmt"
	"math/rand"
)

type SmallIntFlavor struct{}

func (id SmallIntFlavor) Name() string {
	return "small_int"
}

func (id SmallIntFlavor) Generate(e string) string {
	return fmt.Sprintf("%d", rand.Intn(30))
}

func (id SmallIntFlavor) Flavor() string {
	return "int"
}
