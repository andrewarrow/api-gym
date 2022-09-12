package flavor

import (
	"fmt"

	"github.com/brianvoe/gofakeit/v6"
)

type LongitudeFlavor struct{}

func (id LongitudeFlavor) Name() string {
	return "longitude"
}

func (id LongitudeFlavor) Generate(e string) string {
	return fmt.Sprintf("%f", gofakeit.Longitude())
}
func (id LongitudeFlavor) Flavor() string {
	return "float64"
}
