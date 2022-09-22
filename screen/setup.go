package screen

import (
	"fmt"
	"log"
	"math/rand"
	"os"
	"strings"

	"github.com/brianvoe/gofakeit/v6"
	ui "github.com/gizak/termui/v3"
	"github.com/gizak/termui/v3/widgets"
)

var grid = ui.NewGrid()
var flavors = widgets.NewList()
var selected = widgets.NewList()
var rendered = widgets.NewList()
var tab = "flavors"
var insertMode = false

type GymField struct {
	Flavor string
	Name   string
}

var selectedItems = []GymField{}

func Setup() {
	if err := ui.Init(); err != nil {
		log.Fatalf("failed to initialize termui: %v", err)
	}
	defer ui.Close()

	setListColors(flavors)
	setListColors(selected)
	setListColors(rendered)

	flavors.Rows = []string{"individual", "location", "phone", "timestamp", "few_words", "int", "float", "bool", "paragraph"}

	termWidth, termHeight := ui.TerminalDimensions()
	grid.SetRect(0, 0, termWidth, termHeight)

	grid.Set(
		ui.NewRow(1.0,
			ui.NewCol(0.16, flavors),
			ui.NewCol(0.16, selected),
			ui.NewCol(0.68, rendered),
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
	if e.ID == "<Enter>" || e.ID == "<Escape>" {
		insertMode = false
	} else if e.ID == "<Backspace>" {
		i := selected.SelectedRow
		text := selected.Rows[i]
		if len(text) > 0 {
			selected.Rows[i] = text[0 : len(text)-1]
		}
	} else if e.ID == "<Space>" {
		i := selected.SelectedRow
		selected.Rows[i] += "_"
	} else {
		i := selected.SelectedRow
		selected.Rows[i] += strings.ToLower(e.ID)
	}
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
	case "i":
		insertMode = true
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
	if tab == "flavors" {
		handleEnterFlavors()
	} else if tab == "selected" {
		handleEnterSelected()
	}
}

func handleEnterSelected() {
	rendered.Rows = []string{}
	for _, item := range selectedItems {
		val := ""
		if item.Flavor == "address" {
			a := gofakeit.Address()
			val = fmt.Sprintf("%s, %s, %s %s %s", a.Street, a.City, a.State, a.Zip, a.Country)
		} else if item.Flavor == "latitude" {
			val = fmt.Sprintf("%f", gofakeit.Latitude())
		} else if item.Flavor == "longitude" {
			val = fmt.Sprintf("%f", gofakeit.Longitude())
		} else if item.Flavor == "timestamp" {
			val = "2022-04-18T06:52:29.940Z"
		} else if item.Flavor == "few_words" {
			val = gofakeit.Word() + " " + gofakeit.Word()
		} else if item.Flavor == "float" {
			val = fmt.Sprintf("%d.%d", rand.Intn(30), rand.Intn(10))
		} else if item.Flavor == "int" {
			val = fmt.Sprintf("%d", rand.Intn(65000))
		} else if item.Flavor == "paragraph" {
			val = gofakeit.LoremIpsumParagraph(1, 3, 33, ".")
		}
		rendered.Rows = append(rendered.Rows, val)
	}
}

func handleEnterFlavors() {
	if flavors.SelectedRow == 1 {
		addToSelectedItems("address", "address")
		addToSelectedItems("latitude", "latitude")
		addToSelectedItems("longitude", "longitude")
	} else if flavors.SelectedRow == 3 {
		addToSelectedItems("timestamp", "something_at")
	} else if flavors.SelectedRow == 4 {
		addToSelectedItems("few_words", "few_words")
	} else if flavors.SelectedRow == 5 {
		addToSelectedItems("int", "int")
	} else if flavors.SelectedRow == 6 {
		addToSelectedItems("float", "float")
	} else if flavors.SelectedRow == 8 {
		addToSelectedItems("paragraph", "paragraph")
	}
	//selected.SelectedRow = len(selected.Rows) + 1
}

func addToSelectedItems(flavor, name string) {
	gf := GymField{}
	gf.Flavor = flavor
	gf.Name = name
	selectedItems = append(selectedItems, gf)
	selected.Rows = append(selected.Rows, name)
}

func setListColors(s *widgets.List) {
	s.SelectedRowStyle.Fg = ui.ColorWhite
	s.SelectedRowStyle.Bg = ui.ColorMagenta
	s.TextStyle.Fg = ui.ColorWhite
	s.TextStyle.Bg = ui.ColorBlack
}
