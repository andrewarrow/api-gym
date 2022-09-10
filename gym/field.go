package gym

import (
	"api-gym/util"
	"fmt"
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
