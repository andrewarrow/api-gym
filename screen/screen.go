package screen

import (
	"api-gym/gym"
	"fmt"
	"log"

	ui "github.com/gizak/termui/v3"
	"github.com/gizak/termui/v3/widgets"
)

type GymScreen struct {
	models *widgets.List
	fields *widgets.List
	g      *gym.Gym
}

func Run(g *gym.Gym) {
	if err := ui.Init(); err != nil {
		log.Fatalf("failed to initialize termui: %v", err)
	}
	defer ui.Close()

	gymScreen := &GymScreen{}
	gymScreen.g = g
	gymScreen.models = widgets.NewList()
	gymScreen.models.Title = "Models"
	gymScreen.models.Rows = []string{}
	gymScreen.models.SelectedRowStyle.Fg = ui.ColorWhite
	gymScreen.models.SelectedRowStyle.Bg = ui.ColorGreen
	gymScreen.models.TextStyle.Fg = ui.ColorWhite
	gymScreen.models.TextStyle.Bg = ui.ColorBlack
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
		e := <-uiEvents
		switch e.ID {
		case "q", "<C-c>":
			return
		case "j", "<Down>":
			gs.models.ScrollDown()
		case "k", "<Up>":
			gs.models.ScrollUp()
		case "<C-d>":
			gs.models.ScrollHalfPageDown()
		case "<C-u>":
			gs.models.ScrollHalfPageUp()
		case "<C-f>":
			gs.models.ScrollPageDown()
		case "<C-b>":
			gs.models.ScrollPageUp()
		case "<Enter>":
			gs.fields = widgets.NewList()
			gs.fields.Title = "Fields"
			gs.fields.Rows = []string{}
			gs.fields.SelectedRowStyle.Fg = ui.ColorWhite
			gs.fields.SelectedRowStyle.Bg = ui.ColorGreen
			gs.fields.TextStyle.Fg = ui.ColorWhite
			gs.fields.TextStyle.Bg = ui.ColorBlack
			for i, f := range gs.g.Structs[gs.models.SelectedRow].Fields {
				gs.fields.Rows = append(gs.fields.Rows, fmt.Sprintf("[%d] %s", i, f.Name))
			}
			gs.fields.SetRect(40, 0, 30, 8)
			ui.Render(gs.fields)
		case "<Home>":
			gs.models.ScrollTop()
		case "G", "<End>":
			gs.models.ScrollBottom()
		}

		ui.Render(gs.models)
	}
}
