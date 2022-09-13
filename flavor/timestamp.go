package flavor

import (
	"fmt"
	"math/rand"
)

type TimestampFlavor struct {
	Options []string
}

func NewTimestampFlavor() TimestampFlavor {
	f := TimestampFlavor{}
	f.Options = []string{"years_with_null", "years_never_null"}
	return f
}

func (f TimestampFlavor) Name() string {
	return "timestamp"
}

func (f TimestampFlavor) Generate(e string) string {
	value := "2022-04-18T06:52:29.940Z"
	if e == "1" { // years_with_null
		if rand.Intn(2) == 0 {
			value = "null"
		}
	} else if e == "2" { //years_never_null
	}
	return value
}

func (f TimestampFlavor) Flavor() string {
	return "string"
}

func (f TimestampFlavor) ListOptions() bool {
	for i, option := range f.Options {
		fmt.Printf("%2d. %s\n", i+1, option)
	}
	return true
}
