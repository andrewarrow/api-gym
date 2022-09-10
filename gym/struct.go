package gym

import "strings"

type Struct struct {
	Name string `json:"name"`
}

func NewStruct(name string) *Struct {
	s := Struct{}
	s.Name = name
	if strings.HasPrefix(name, "[]") {
		s.Name = name[2:]
	}
	return &s
}
