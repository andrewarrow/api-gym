package flavor

import (
	"fmt"
	"math/rand"
)

type SmallFloatFlavor struct{}

func (id SmallFloatFlavor) Name() string {
	return "small_float"
}

func (id SmallFloatFlavor) Generate() string {
	return fmt.Sprintf("%d.%d", rand.Intn(30), rand.Intn(10))
}
