package screen

import (
	"api-gym/gym"
	"fmt"
	"log"

	ui "github.com/gizak/termui/v3"
	"github.com/gizak/termui/v3/widgets"
)

type GymScreen struct {
	activeListName string
	listMap        map[string]*widgets.List
	listArray      []string
	listIndex      int
	g              *gym.Gym
}

func Run(g *gym.Gym) {
	if err := ui.Init(); err != nil {
		log.Fatalf("failed to initialize termui: %v", err)
	}
	defer ui.Close()

	gs := &GymScreen{}
	gs.listMap = map[string]*widgets.List{}
	gs.g = g
	models := MakeNewList("Models", ui.ColorMagenta)
	gs.listMap["models"] = models
	gs.listArray = append(gs.listArray, "models")
	gs.listIndex = 0
	gs.activeListName = "models"
	for i, s := range g.Structs {
		models.Rows = append(models.Rows, fmt.Sprintf("[%d] %s", i, s.Name))
	}
	models.SetRect(0, 0, 26, 8)

	fields := MakeNewList("Fields", ui.ColorGreen)
	fields.SetRect(27, 0, 56+27, 8)
	gs.listMap["fields"] = fields
	gs.listArray = append(gs.listArray, "fields")

	list := MakeNewList("Flavors", ui.ColorGreen)
	list.Rows = append(list.Rows, "Address")
	list.Rows = append(list.Rows, "Timestamp")
	list.Rows = append(list.Rows, "LargeInt")
	list.Rows = append(list.Rows, "SmallInt")
	list.Rows = append(list.Rows, "[]OtherModel")
	list.SetRect(27, 8, 26+27, 8+8)
	gs.listMap["flavors"] = list
	gs.listArray = append(gs.listArray, "flavors")

	ui.Render(gs.activeLists()...)
	gs.poll()
}

func (gs *GymScreen) poll() {

	uiEvents := ui.PollEvents()
	for {
		e := <-uiEvents
		switch e.ID {
		case "q", "<C-c>":
			return
		case "j", "<Down>":
			gs.activeList().ScrollDown()
		case "k", "<Up>":
			gs.activeList().ScrollUp()
		case "a":
			gs.addNew()
		case "<Escape>":
		case "<Tab>":
			gs.nextList()
		case "<Enter>":
			if gs.activeListName == "models" {
				gs.enterOnModels()
			} else if gs.activeListName == "fields" {
				gs.enterOnFields()
			} else if gs.activeListName == "flavors" {
			}
		}

		ui.Render(gs.activeLists()...)
	}
}

func (gs *GymScreen) nextList() {
	gs.listIndex++
	if gs.listIndex >= len(gs.listArray) {
		gs.listIndex = 0
	}
	gs.setActiveByName(gs.listArray[gs.listIndex])
}

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

func (gs *GymScreen) enterOnFields() {
	gs.setActiveByName("flavors")
}

func (gs *GymScreen) setActiveByName(name string) {
	gs.activeListName = name
	for i, listName := range gs.listArray {
		if listName == name {
			gs.listIndex = i
			break
		}
	}
	for _, listName := range gs.listArray {
		list := gs.listMap[listName]
		list.SelectedRowStyle.Bg = ui.ColorGreen
	}
	gs.activeListName = gs.listArray[gs.listIndex]
	list := gs.listMap[gs.activeListName]
	list.SelectedRowStyle.Bg = ui.ColorMagenta
}

func (gs *GymScreen) addNew() {
	if gs.activeListName == "fields" {
		//s := gs.g.Structs[gs.models.SelectedRow]
		list := gs.listMap["fields"]
		list.Rows = append(list.Rows, fmt.Sprintf("[%d] %s", len(list.Rows), "Untitled"))
	}
}

func (gs *GymScreen) activeList() *widgets.List {
	return gs.listMap[gs.activeListName]
}

func (gs *GymScreen) activeLists() []ui.Drawable {
	items := []ui.Drawable{}
	for _, v := range gs.listMap {
		if v == nil {
			continue
		}
		items = append(items, v)
	}
	return items
}

func MakeNewList(title string, color ui.Color) *widgets.List {
	list := widgets.NewList()
	list.Title = title
	list.Rows = []string{}
	list.SelectedRowStyle.Fg = ui.ColorWhite
	list.SelectedRowStyle.Bg = color
	list.TextStyle.Fg = ui.ColorWhite
	list.TextStyle.Bg = ui.ColorBlack
	return list
}
