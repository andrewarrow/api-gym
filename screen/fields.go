package screen

import "api-gym/gym"

func (gs *GymScreen) enterOnFields() {
	gs.setActiveByName("flavors")
}

func (gs *GymScreen) deleteField() {
	models := gs.listMap["models"]
	fields := gs.listMap["fields"]
	newList := []*gym.Field{}
	for i, f := range gs.g.Structs[models.SelectedRow].Fields {
		if i == fields.SelectedRow {
			continue
		}
		newList = append(newList, f)
	}
	gs.g.Structs[models.SelectedRow].Fields = newList
	gs.enterOnModels()
}
