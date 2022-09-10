package gym

type Field struct {
	Name   string `json:"name"`
	Flavor string `json:"flavor"`
}

func NewField(name, flavor string) *Field {
	f := Field{}
	f.Name = name
	f.Flavor = flavor
	return &f
}
