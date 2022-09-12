package gym

import (
	"api-gym/util"
	"strings"
)

type Struct struct {
	Name       string   `json:"name"`
	Fields     []*Field `json:"field"`
	ArrayOrMap string   `json:"array_or_map"`
}

func NewStruct(name string) *Struct {
	s := Struct{}
	s.Name = name
	s.ArrayOrMap = "array"
	if strings.HasPrefix(name, "[]") {
		s.Name = name[2:]
	}
	s.Fields = append(s.Fields, NewField("Id", "string", "uuid"))
	s.Fields = append(s.Fields, NewField("Name", "string", "few_words"))
	return &s
}

func (s *Struct) JsonContainerName() string {
	return util.ToSnakeCase(s.Name) + "s"
}
