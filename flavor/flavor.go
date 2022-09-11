package flavor

import "fmt"

type Flavor interface {
	String() string
}

var allFlavors = []Flavor{IdFlavor{}, NameFlavor{}}

func ListFlavors() {
	fmt.Println("")
	for i, flavor := range allFlavors {
		fmt.Printf("%2d. %-30s\n", i, flavor.String())
	}
	fmt.Println("")
}

type IdFlavor struct {
}

func (id IdFlavor) String() string {
	return "uuid"
}

type NameFlavor struct {
}

func (id NameFlavor) String() string {
	return "name"
}
