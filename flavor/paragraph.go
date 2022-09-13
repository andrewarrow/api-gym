package flavor

import "github.com/brianvoe/gofakeit/v6"

type ParagraphFlavor struct{}

func (id ParagraphFlavor) Name() string {
	return "paragraph"
}

func (id ParagraphFlavor) Generate(e string) string {
	return gofakeit.LoremIpsumParagraph(1, 3, 33, ".")
}
func (id ParagraphFlavor) Flavor() string {
	return "string"
}
func (f ParagraphFlavor) ListOptions() bool {
	return false
}
