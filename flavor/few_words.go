package flavor

import "github.com/brianvoe/gofakeit/v6"

type FewWordsFlavor struct{}

func (id FewWordsFlavor) Name() string {
	return "few_words"
}

func (id FewWordsFlavor) Generate() string {
	return gofakeit.Word() + " " + gofakeit.Word()
}
