package flavor

import (
	"api-gym/util"
	"fmt"

	"github.com/brianvoe/gofakeit/v6"
)

type Flavor interface {
	Name() string
	Generate() string
}

var allFlavors = []Flavor{IdFlavor{}, NameFlavor{}, FewWordsFlavor{},
	AddressFlavor{},
	LatitudeFlavor{},
	LongitudeFlavor{},
	PronounsFlavor{},
	LargeIntFlavor{},
	SmallFloatFlavor{},
	SmallIntFlavor{},
	TimestampFlavor{},
	ParagraphFlavor{}}

func FlavorsAsMap() map[string]Flavor {
	m := map[string]Flavor{}
	for _, flavor := range allFlavors {
		m[flavor.Name()] = flavor
	}
	return m
}

func ListFlavors() {
	fmt.Println("")
	for i, flavor := range allFlavors {
		fmt.Printf("%2d. %-30s %-30s\n", i, flavor.Name(), flavor.Generate())
	}
	fmt.Println("")
}

type IdFlavor struct {
}

func (id IdFlavor) Name() string {
	return "uuid"
}

func (id IdFlavor) Generate() string {
	return util.PseudoUuid()
}

type NameFlavor struct {
}

func (id NameFlavor) Name() string {
	return "name"
}

func (id NameFlavor) Generate() string {
	return gofakeit.Name()
}
