package flavor

import "math/rand"

type PronounsFlavor struct{}

func (id PronounsFlavor) Name() string {
	return "pronouns"
}

func (id PronounsFlavor) Generate() string {
	value := "she/her"
	if rand.Intn(2) == 0 {
		value = "he/him"
	}
	if rand.Intn(2) == 0 {
		value = "they/them"
	}
	return value
}
func (id PronounsFlavor) Flavor() string {
	return "string"
}
