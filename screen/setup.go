package screen

import (
	"log"

	ui "github.com/gizak/termui/v3"
	"github.com/gizak/termui/v3/widgets"
)

var flavors = widgets.NewList()

func Setup() {
	if err := ui.Init(); err != nil {
		log.Fatalf("failed to initialize termui: %v", err)
	}
	defer ui.Close()

	flavors.SelectedRowStyle.Fg = ui.ColorWhite
	flavors.SelectedRowStyle.Bg = ui.ColorMagenta
	flavors.TextStyle.Fg = ui.ColorWhite
	flavors.TextStyle.Bg = ui.ColorBlack
	flavors.Rows = []string{"individual", "location", "timestamp", "few_words", "int", "float", "paragraph"}

	grid := ui.NewGrid()
	termWidth, termHeight := ui.TerminalDimensions()
	grid.SetRect(0, 0, termWidth, termHeight)

	grid.Set(
		ui.NewRow(1.0,
			ui.NewCol(0.50, flavors),
		),
	)

	ui.Render(grid)
	uiEvents := ui.PollEvents()
	for {
		select {
		case e := <-uiEvents:
			switch e.ID {
			case "q", "<C-c>":
				return
			case "j", "<Down>":
			case "k", "<Up>":
			case "<Tab>":
			case "<Enter>":
			case "<Resize>":
				payload := e.Payload.(ui.Resize)
				grid.SetRect(0, 0, payload.Width, payload.Height)
				ui.Clear()
			}
		}
		ui.Render(grid)
	}
}
