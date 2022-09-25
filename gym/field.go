package gym

import (
	"api-gym/flavor"
	"api-gym/util"
	"fmt"
)

type Field struct {
	Name   string `json:"name"`
	Flavor string `json:"flavor"`
	Extra  string `json:"extra"`
}

func NewField(name, flavor, extra string) *Field {
	f := Field{}
	f.Name = name
	f.Flavor = flavor
	f.Extra = extra
	return &f
}

func (f *Field) DataType() string {
	return flavor.DataType(f.Flavor, f.Extra)
}
func (f *Field) NameToJson() string {
	return fmt.Sprintf("%s", util.ToSnakeCase(f.Name))
}

func (f *Field) FlavorToStructName() string {
	return f.Extra
}

func (f *Field) ToFakeValue() string {
	return flavor.Generate(f.Flavor, f.Extra)
}
