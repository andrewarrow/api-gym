package screen

import (
	"api-gym/gym"
	"strings"
)

func (gs *GymScreen) enterOnFields() {
	gs.setActiveByName("flavors")
}

func (gs *GymScreen) deleteField() {
	models := gs.listMap["models"]
	fields := gs.listMap["fields"]
	gs.g.DeleteField(models.SelectedRow, fields.SelectedRow)
	gs.enterOnModels()
}

func (gs *GymScreen) renameField() {
	field := gs.fieldFromModel()
	field.Name = ""
	gs.enterOnModels()
	gs.rename = true
}

func (gs *GymScreen) renameFieldEvent(c string) {
	field := gs.fieldFromModel()

	if len(field.Name) == 0 {
		field.Name += strings.ToUpper(c)
	} else {
		field.Name += c
	}
	gs.enterOnModels()
}

func (gs *GymScreen) backspaceFieldEvent() {
	field := gs.fieldFromModel()

	name := field.Name
	field.Name = name[0 : len(name)-1]
	gs.enterOnModels()
}

func (gs *GymScreen) fieldFromModel() *gym.Field {
	models := gs.listMap["models"]
	fields := gs.listMap["fields"]
	model := gs.g.Structs[models.SelectedRow]
	field := model.Fields[fields.SelectedRow]
	return field
}
