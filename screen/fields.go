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
