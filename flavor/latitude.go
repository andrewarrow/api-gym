package flavor

import (
	"fmt"

	"github.com/brianvoe/gofakeit/v6"
)

type LatitudeFlavor struct{}

func (id LatitudeFlavor) Name() string {
	return "latitude"
}

func (id LatitudeFlavor) Generate(e string) string {
	return fmt.Sprintf("%f", gofakeit.Latitude())
}

func (id LatitudeFlavor) Flavor() string {
	return "float64"
}
