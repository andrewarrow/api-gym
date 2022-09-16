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
	g              *gym.Gym
	quit           bool
}

func Run(g *gym.Gym) {
	if err := ui.Init(); err != nil {
		log.Fatalf("failed to initialize termui: %v", err)
	}
	defer ui.Close()

	gymScreen := &GymScreen{}
	gymScreen.listMap = map[string]*widgets.List{}
	gymScreen.g = g
	models := MakeNewList("Models")
	gymScreen.listMap["models"] = models
	gymScreen.activeListName = "models"
	for i, s := range g.Structs {
		models.Rows = append(models.Rows, fmt.Sprintf("[%d] %s", i, s.Name))
	}
	models.SetRect(0, 0, 30, 8)
	ui.Render(models)
	gymScreen.poll()
}

func (gs *GymScreen) poll() {

	uiEvents := ui.PollEvents()
	for {
		if gs.quit {
			break
		}
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
			gs.listMap["fields"] = nil
			gs.activeListName = "models"
			gs.quit = true
		case "<Enter>":
			if gs.activeListName == "models" {
				gs.enterOnModels()
			} else if gs.activeListName == "fields" {
				gs.enterOnFields()
			}
		}

		ui.Render(gs.activeLists()...)
	}

	Run(gs.g)
}

func (gs *GymScreen) enterOnModels() {
	list := MakeNewList("Fields")
	models := gs.listMap["models"]
	for i, f := range gs.g.Structs[models.SelectedRow].Fields {
		list.Rows = append(list.Rows, fmt.Sprintf("[%02d] %-10s %s", i, f.Name, f.Random))
	}
	list.Rows = append(list.Rows, "     ADD NEW")
	list.SetRect(31, 0, 30+31, 8)
	gs.listMap["fields"] = list
	gs.activeListName = "fields"
}

func (gs *GymScreen) enterOnFields() {
	list := MakeNewList("Flavors")
	list.Rows = append(list.Rows, "Address")
	list.Rows = append(list.Rows, "Timestamp")
	list.Rows = append(list.Rows, "LargeInt")
	list.Rows = append(list.Rows, "SmallInt")
	list.Rows = append(list.Rows, "[]OtherModel")
	list.SetRect(31, 9, 30+31, 9+8)
	gs.listMap["flavors"] = list
	gs.activeListName = "flavors"
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

func MakeNewList(title string) *widgets.List {
	list := widgets.NewList()
	list.Title = title
	list.Rows = []string{}
	list.SelectedRowStyle.Fg = ui.ColorWhite
	list.SelectedRowStyle.Bg = ui.ColorGreen
	list.TextStyle.Fg = ui.ColorWhite
	list.TextStyle.Bg = ui.ColorBlack
	return list
}
