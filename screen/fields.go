package screen

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
	gs.g.Structs[models.SelectedRow].Fields[fields.SelectedRow].Name = ""
	gs.enterOnModels()
	gs.rename = true
}

func (gs *GymScreen) renameFieldEvent(c string) {
	models := gs.listMap["models"]
	fields := gs.listMap["fields"]
	gs.g.Structs[models.SelectedRow].Fields[fields.SelectedRow].Name += c
	gs.enterOnModels()
}
