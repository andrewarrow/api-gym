package flavor

import "math/rand"

type BooleanFlavor struct{}

func (id BooleanFlavor) Name() string {
	return "boolean"
}

func (id BooleanFlavor) Generate(e string) string {
	if rand.Intn(2) == 0 {
		return "false"
	}
	return "true"
}

func (id BooleanFlavor) Flavor() string {
	return "bool"
}
func (f BooleanFlavor) ListOptions() bool {
	return false
}
