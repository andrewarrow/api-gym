package screen

func (gs *GymScreen) enterOnFlavors() {
	models := gs.listMap["models"]
	//fields := gs.listMap["fields"]
	flavors := gs.listMap["flavors"]

	gs.g.AddFieldToStruct(models.SelectedRow, flavors.SelectedRow)
	gs.enterOnModels()
}

func flavorList() []string {
	items := []string{}
	items = append(items, "uuid")
	items = append(items, "name")
	items = append(items, "few_words")
	items = append(items, "address")
	items = append(items, "pronouns")
	items = append(items, "int")
	items = append(items, "float")
	items = append(items, "timestamp")
	items = append(items, "enum")
	items = append(items, "bool")
	items = append(items, "paragraph")
	return items
}
