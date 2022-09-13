package flavor

import (
	"fmt"
	"math/rand"
)

type SmallFloatFlavor struct{}

func (id SmallFloatFlavor) Name() string {
	return "small_float"
}

func (id SmallFloatFlavor) Generate(e string) string {
	return fmt.Sprintf("%d.%d", rand.Intn(30), rand.Intn(10))
}

func (id SmallFloatFlavor) Flavor() string {
	return "float64"
}
func (f SmallFloatFlavor) ListOptions() bool {
	return false
}
