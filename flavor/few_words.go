package flavor

import "github.com/brianvoe/gofakeit/v6"

type FewWordsFlavor struct{}

func (id FewWordsFlavor) Name() string {
	return "few_words"
}

func (id FewWordsFlavor) Generate() string {
	value := gofakeit.Word() + " " + gofakeit.Word()
	for {
		if len(value) >= 33 {
			break
		}
		value = value + " " + gofakeit.Word()
	}
	return value
}
