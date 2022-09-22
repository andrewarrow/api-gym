package screen

import (
	"fmt"
	"log"
	"os"

	ui "github.com/gizak/termui/v3"
	"github.com/gizak/termui/v3/widgets"
)

var grid = ui.NewGrid()
var flavors = widgets.NewList()
var selected = widgets.NewList()
var tab = "flavors"
var insertMode = false

func Setup() {
	if err := ui.Init(); err != nil {
		log.Fatalf("failed to initialize termui: %v", err)
	}
	defer ui.Close()

	setListColors(flavors)
	flavors.Rows = []string{"individual", "location", "phone", "timestamp", "few_words", "int", "float", "bool", "paragraph"}
	setListColors(selected)

	termWidth, termHeight := ui.TerminalDimensions()
	grid.SetRect(0, 0, termWidth, termHeight)

	grid.Set(
		ui.NewRow(1.0,
			ui.NewCol(0.16, flavors),
			ui.NewCol(0.16, selected),
		),
	)

	ui.Render(grid)
	uiEvents := ui.PollEvents()
	for {
		e := <-uiEvents
		if insertMode {
			handleInsert(e)
		} else {
			normalEvents(e)
		}
		ui.Render(grid)
	}
}

func handleInsert(e ui.Event) {
}

func normalEvents(e ui.Event) {
	switch e.ID {
	case "q", "<C-c>":
		ui.Close()
		fmt.Println("")
		fmt.Println("")
		fmt.Println("")
		fmt.Println("")
		os.Exit(1)
		return
	case "j", "<Down>":
		selectedList().ScrollDown()
	case "k", "<Up>":
		selectedList().ScrollUp()
	case "<Right>":
		if tab == "flavors" {
			tab = "selected"
		} else if tab == "selected" {
			tab = "flavors"
		}
	case "<Left>":
		if tab == "flavors" {
			tab = "selected"
		} else if tab == "selected" {
			tab = "flavors"
		}
	case "<Tab>":
		if tab == "flavors" {
			tab = "selected"
		} else if tab == "selected" {
			tab = "flavors"
		}
	case "<Enter>":
		handleEnter()
	case "<Resize>":
		payload := e.Payload.(ui.Resize)
		grid.SetRect(0, 0, payload.Width, payload.Height)
		ui.Clear()
	}
}

func selectedList() *widgets.List {
	if tab == "flavors" {
		return flavors
	} else if tab == "selected" {
		return selected
	}

	return selected
}

func handleEnter() {
	if flavors.SelectedRow == 1 {
		selected.Rows = append(selected.Rows, "address")
		selected.Rows = append(selected.Rows, "latitude")
		selected.Rows = append(selected.Rows, "longitude")
	} else if flavors.SelectedRow == 3 {
		selected.Rows = append(selected.Rows, "something_at")
	} else if flavors.SelectedRow == 4 {
		selected.Rows = append(selected.Rows, "few_words")
	} else if flavors.SelectedRow == 5 {
		selected.Rows = append(selected.Rows, "int")
	} else if flavors.SelectedRow == 8 {
		selected.Rows = append(selected.Rows, "paragraph")
	}
	//selected.SelectedRow = len(selected.Rows) + 1
}

func setListColors(s *widgets.List) {
	s.SelectedRowStyle.Fg = ui.ColorWhite
	s.SelectedRowStyle.Bg = ui.ColorMagenta
	s.TextStyle.Fg = ui.ColorWhite
	s.TextStyle.Bg = ui.ColorBlack
}
