package screen

func (gs *GymScreen) enterOnFlavors() {
	models := gs.listMap["models"]
	//fields := gs.listMap["fields"]
	flavors := gs.listMap["flavors"]

	gs.g.AddFieldToStruct(models.SelectedRow, flavors.SelectedRow)
	gs.enterOnModels()
}
