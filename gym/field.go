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
	Extra  string `json:"extra"`
}

func NewField(name, flavor, random, extra string) *Field {
	f := Field{}
	f.Name = name
	f.Flavor = flavor
	f.Random = random
	f.Extra = extra
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
	return theFlavor.Generate(f.Extra)
}
