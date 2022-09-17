package screen

import "strings"

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
	models := gs.listMap["models"]
	fields := gs.listMap["fields"]
	model := gs.g.Structs[models.SelectedRow]
	field := model.Fields[fields.SelectedRow]
	field.Name = ""
	gs.enterOnModels()
	gs.rename = true
}

func (gs *GymScreen) renameFieldEvent(c string) {
	models := gs.listMap["models"]
	fields := gs.listMap["fields"]
	model := gs.g.Structs[models.SelectedRow]
	field := model.Fields[fields.SelectedRow]

	if len(field.Name) == 0 {
		field.Name += strings.ToUpper(c)
	} else {
		field.Name += c
	}
	gs.enterOnModels()
}

func (gs *GymScreen) backspaceFieldEvent() {
	models := gs.listMap["models"]
	fields := gs.listMap["fields"]
	model := gs.g.Structs[models.SelectedRow]
	field := model.Fields[fields.SelectedRow]

	name := field.Name
	field.Name = name[0 : len(name)-1]
	gs.enterOnModels()
}
