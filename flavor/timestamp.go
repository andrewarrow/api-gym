package flavor

import "math/rand"

type TimestampFlavor struct{}

func (id TimestampFlavor) Name() string {
	return "timestamp"
}

func (id TimestampFlavor) Generate(e string) string {
	value := "2022-04-18T06:52:29.940Z"
	if e == "years_with_null" {
		if rand.Intn(2) == 0 {
			value = "null"
		}
	}
	return value
}

func (id TimestampFlavor) Flavor() string {
	return "string"
}
