package gym

import "fmt"

func (g *Gym) ListRoutesAndModels() {
	for i, route := range g.Routes {
		fmt.Printf("%2d. %-30s %s,%d\n", i+1, route.String(), route.Model, route.Count)
	}
	fmt.Println("")
	for i, s := range g.Structs {
		fmt.Printf("%2d. %-30s\n", i+1, s.Name)
		for j, f := range s.Fields {
			fmt.Printf("  %2d. %-20s %-26s  %s\n", j+1, f.Name, f.Flavor, f.Extra)
		}
	}
}
