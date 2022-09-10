package gym

import (
	"api-gym/util"
	"fmt"
	"math/rand"
	"strings"

	"github.com/brianvoe/gofakeit/v6"
)

type Field struct {
	Name   string `json:"name"`
	Flavor string `json:"flavor"`
	Random string `json:"random"`
}

func NewField(name, flavor, random string) *Field {
	f := Field{}
	f.Name = name
	f.Flavor = flavor
	f.Random = random
	return &f
}

func (f *Field) NameToJson() string {
	return fmt.Sprintf("%s", util.ToSnakeCase(f.Name))
}
func (f *Field) FlavorToStructName() string {
	if strings.HasPrefix(f.Flavor, "[]") {
		return f.Flavor[2:]
	}
	return fmt.Sprintf("%s", f.Flavor)
}

func (f *Field) ToFakeValue() string {
	value := ""
	if f.Random == "uuid" {
		value = util.PseudoUuid()
	} else if f.Random == "color" {
		value = gofakeit.Color()
	} else if f.Random == "name" {
		value = gofakeit.Name()
	} else if f.Random == "address" {
		a := gofakeit.Address()
		value = fmt.Sprintf("%s, %s, %s %s %s", a.Street, a.City, a.State, a.Zip, a.Country)
	} else if f.Random == "latitude" {
		value = fmt.Sprintf("%f", gofakeit.Latitude())
	} else if f.Random == "longitude" {
		value = fmt.Sprintf("%f", gofakeit.Longitude())
	} else if f.Random == "small_int_range" {
		value = fmt.Sprintf("%d", rand.Intn(30))
	}
	return value
}
