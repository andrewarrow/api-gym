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
	} else if f.Random == "two_words" {
		value = gofakeit.HipsterWord() + " " + gofakeit.HipsterWord()
	} else if f.Random == "name" {
		value = gofakeit.Name()
	} else if f.Random == "paragraph" {
		value = gofakeit.LoremIpsumParagraph(1, 3, 33, ".")
	} else if f.Random == "pronouns" {
		value = "she/her"
		if rand.Intn(2) == 0 {
			value = "he/him"
		}
		if rand.Intn(2) == 0 {
			value = "they/them"
		}
	} else if f.Random == "address" {
		a := gofakeit.Address()
		value = fmt.Sprintf("%s, %s, %s %s %s", a.Street, a.City, a.State, a.Zip, a.Country)
	} else if f.Random == "latitude" {
		value = fmt.Sprintf("%f", gofakeit.Latitude())
	} else if f.Random == "longitude" {
		value = fmt.Sprintf("%f", gofakeit.Longitude())
	} else if f.Random == "small_int" {
		value = fmt.Sprintf("%d", rand.Intn(30))
	} else if f.Random == "small_float" {
		value = fmt.Sprintf("%d.%d", rand.Intn(30), rand.Intn(10))
	} else if f.Random == "large_int" {
		value = fmt.Sprintf("%d", 6000+rand.Intn(9000))
	} else if f.Random == "timestamp" {
		value = "2022-04-18T06:52:29.940Z"
	}
	return value
}
