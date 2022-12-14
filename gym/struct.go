package gym

import (
	"api-gym/util"
	"strings"
)

type Struct struct {
	Name       string   `json:"name"`
	Fields     []*Field `json:"field"`
	ArrayOrMap string   `json:"array_or_map"`
	Extra      string   `json:"extra"`
}

func NewStruct(name string) *Struct {
	s := Struct{}
	s.Name = name
	s.ArrayOrMap = "array"
	if strings.HasPrefix(name, "[]") {
		s.Name = name[2:]
	}
	return &s
}

func (s *Struct) JsonContainerName() string {
	return util.Plural(s.Name)
}
