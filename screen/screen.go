package screen

import (
	"api-gym/gym"
	"fmt"
	"log"

	ui "github.com/gizak/termui/v3"
	"github.com/gizak/termui/v3/widgets"
)

type GymScreen struct {
	models         *widgets.List
	fields         *widgets.List
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
	gymScreen.models = MakeNewList("Models")
	gymScreen.listMap["models"] = gymScreen.models
	gymScreen.activeListName = "models"
	for i, s := range g.Structs {
		gymScreen.models.Rows = append(gymScreen.models.Rows, fmt.Sprintf("[%d] %s", i, s.Name))
	}
	gymScreen.models.SetRect(0, 0, 30, 8)
	ui.Render(gymScreen.models)
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
			list := MakeNewList("Fields")
			for i, f := range gs.g.Structs[gs.models.SelectedRow].Fields {
				list.Rows = append(list.Rows, fmt.Sprintf("[%d] %s", i, f.Name))
			}
			list.SetRect(40, 0, 60, 8)
			gs.fields = list
			gs.listMap["fields"] = list
			gs.activeListName = "fields"
		}

		ui.Render(gs.activeLists()...)
	}

	Run(gs.g)
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
	items := []ui.Drawable{gs.models}
	if gs.fields != nil {
		items = append(items, gs.fields)
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
