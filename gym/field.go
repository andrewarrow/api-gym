package gym

import (
	"api-gym/util"
	"fmt"
	"math/rand"

	"github.com/brianvoe/gofakeit/v6"
)

type Field struct {
	Name   string `json:"name"`
	Flavor string `json:"flavor"`
}

func NewField(name, flavor string) *Field {
	f := Field{}
	f.Name = name
	f.Flavor = flavor
	return &f
}

func (f *Field) NameToJson() string {
	return fmt.Sprintf("%s", util.ToSnakeCase(f.Name))
}
func (f *Field) ToFakeValue() string {
	value := gofakeit.Color()
	if f.Name == "Id" {
		value = util.PseudoUuid()
	} else if f.Name == "Address" {
		a := gofakeit.Address()
		value = fmt.Sprintf("%s, %s, %s %s %s", a.Street, a.City, a.State, a.Zip, a.Country)
	} else if f.Name == "Latitude" {
		value = fmt.Sprintf("%f", gofakeit.Latitude())
	} else if f.Name == "Longitude" {
		value = fmt.Sprintf("%f", gofakeit.Longitude())
	} else if f.Flavor == "int" {
		value = fmt.Sprintf("%d", rand.Intn(30))
	}
	return value
}
