package flavor

import (
	"fmt"

	"github.com/brianvoe/gofakeit/v6"
)

type AddressFlavor struct{}

func (id AddressFlavor) Name() string {
	return "address"
}

func (id AddressFlavor) Generate() string {
	a := gofakeit.Address()
	return fmt.Sprintf("%s, %s, %s %s %s", a.Street, a.City, a.State, a.Zip, a.Country)
}
func (id AddressFlavor) Flavor() string {
	return "string"
}
