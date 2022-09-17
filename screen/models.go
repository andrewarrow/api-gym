package screen

import "fmt"

func (gs *GymScreen) enterOnModels() {
	models := gs.listMap["models"]
	fields := gs.listMap["fields"]
	fields.Rows = []string{}
	fields.SelectedRow = 0
	for i, f := range gs.g.Structs[models.SelectedRow].Fields {
		fields.Rows = append(fields.Rows, fmt.Sprintf("[%02d] %-16s %-16s %s", i, f.Name, f.Flavor, f.Random))
	}
	fields.Rows = append(fields.Rows, fmt.Sprintf("%50s", ""))
	gs.setActiveByName("fields")
}
