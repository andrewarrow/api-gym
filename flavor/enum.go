package flavor

import (
	"math/rand"
	"strings"
)

type EnumFlavor struct{}

func (id EnumFlavor) Name() string {
	return "enum"
}

func (id EnumFlavor) Generate(extra string) string {
	if extra == "" {
		return "enum"
	}
	tokens := strings.Split(extra, ",")
	index := rand.Intn(len(tokens))
	return tokens[index]
}

func (id EnumFlavor) Flavor() string {
	return "string"
}
func (f EnumFlavor) ListOptions() bool {
	return false
}
