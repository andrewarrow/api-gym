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

func (g *GymScreen) poll() {

	previousKey := ""
	uiEvents := ui.PollEvents()
	for {
		e := <-uiEvents
		switch e.ID {
		case "q", "<C-c>":
			return
		case "j", "<Down>":
			g.models.ScrollDown()
		case "k", "<Up>":
			g.models.ScrollUp()
		case "<C-d>":
			g.models.ScrollHalfPageDown()
		case "<C-u>":
			g.models.ScrollHalfPageUp()
		case "<C-f>":
			g.models.ScrollPageDown()
		case "<C-b>":
			g.models.ScrollPageUp()
		case "<Enter>":
			g.fields = widgets.NewList()
			g.fields.Title = "Fields"
			g.fields.Rows = []string{}
			g.fields.SelectedRowStyle.Fg = ui.ColorWhite
			g.fields.SelectedRowStyle.Bg = ui.ColorGreen
			g.fields.TextStyle.Fg = ui.ColorWhite
			g.fields.TextStyle.Bg = ui.ColorBlack
			for i, f := range g.g.Structs[g.models.SelectedRow].Fields {
				g.fields.Rows = append(g.fields.Rows, fmt.Sprintf("[%d] %s", i, f.Name))
			}
			g.fields.SetRect(40, 0, 30, 8)
			ui.Render(g.fields)
		case "g":
			if previousKey == "g" {
				g.models.ScrollTop()
			}
		case "<Home>":
			g.models.ScrollTop()
		case "G", "<End>":
			g.models.ScrollBottom()
		}

		if previousKey == "g" {
			previousKey = ""
		} else {
			previousKey = e.ID
		}

		ui.Render(g.models)
	}
}
