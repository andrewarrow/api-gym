package screen

import (
	"log"
	"os"

	ui "github.com/gizak/termui/v3"
	"github.com/gizak/termui/v3/widgets"
)

type GymScreen struct {
	models *widgets.List
}

func Run() {
	if err := ui.Init(); err != nil {
		log.Fatalf("failed to initialize termui: %v", err)
	}
	defer ui.Close()

	gymScreen := &GymScreen{}
	gymScreen.models = widgets.NewList()
	gymScreen.models.Title = "Models"
	gymScreen.models.Rows = []string{
		"[0] github.com/gizak/termui/v3",
		"[1] [你好，世界](fg:blue)",
		"[2] [こんにちは世界](fg:red)",
		"[3] [color](fg:white,bg:green) output",
		"[4] output.go",
		"[5] random_out.go",
		"[6] dashboard.go",
		"[7] foo",
		"[8] bar",
		"[9] baz",
	}
	//l.TextStyle = ui.NewStyle(ui.ColorYellow)
	//l.WrapText = false
	gymScreen.models.SetRect(0, 0, 60, 8)
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
			os.Exit(1)
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
