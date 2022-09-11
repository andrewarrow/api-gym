package gym

import (
	"api-gym/flavor"
	"api-gym/util"
	"fmt"
	"strings"
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
	theFlavor := flavor.FlavorsAsMap()[f.Random]
	return theFlavor.Generate()
}
