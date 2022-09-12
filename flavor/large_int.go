package flavor

import (
	"fmt"
	"math/rand"
)

type LargeIntFlavor struct{}

func (id LargeIntFlavor) Name() string {
	return "large_int"
}

func (id LargeIntFlavor) Generate(e string) string {
	return fmt.Sprintf("%d", 6000+rand.Intn(9000))
}

func (id LargeIntFlavor) Flavor() string {
	return "int64"
}
