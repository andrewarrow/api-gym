package gym

import "strings"

type Struct struct {
	Name   string   `json:"name"`
	Fields []*Field `json:"field"`
}

func NewStruct(name string) *Struct {
	s := Struct{}
	s.Name = name
	if strings.HasPrefix(name, "[]") {
		s.Name = name[2:]
	}
	s.Fields = append(s.Fields, NewField("Id", "string", "uuid"))
	s.Fields = append(s.Fields, NewField("Name", "string", "two_words"))
	return &s
}
