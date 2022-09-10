package gym

import "api-gym/util"

type Gym struct {
	Name string
}

func NewGym() *Gym {
	g := Gym{}
	g.Name = util.PseudoUuid()
	return &g
}
