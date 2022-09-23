package screen

import (
	"api-gym/flavor"
	"api-gym/gym"
	"fmt"
	"log"
	"os"
	"strings"

	ui "github.com/gizak/termui/v3"
	"github.com/gizak/termui/v3/widgets"
)

var grid = ui.NewGrid()
var flavors = widgets.NewList()
var selected = widgets.NewList()
var rendered = widgets.NewList()
var tab = "flavors"
var insertMode = false
var selectedItems = []GymField{}

type GymField struct {
	Flavor string
	Name   string
	Extra  string
}

type GymModel struct {
	Name string
	Gym  *gym.Gym
}

func EditModel(model string, g *gym.Gym) {
	s := g.StructsByName[model]
	for _, field := range s.Fields {
		addToSelectedItems(field.Flavor, field.Name, field.Extra)
	}
	StartModelUI(model, g)
}
func AddModel(model string, g *gym.Gym) {
	addToSelectedItems("uuid", "id", "")
	StartModelUI(model, g)
}

func StartModelUI(model string, g *gym.Gym) {
	if err := ui.Init(); err != nil {
		log.Fatalf("failed to initialize termui: %v", err)
	}
	defer ui.Close()
	gm := GymModel{}
	gm.Name = model
	gm.Gym = g

	setListColors(flavors)
	setListColors(selected)
	setListColors(rendered)

	flavors.Rows = []string{"individual", "location", "phone", "timestamp", "few_words",
		"int", "float", "bool", "paragraph", "email", "uuid", "model", "[]model"}

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
			gm.normalEvents(e)
		}
		ui.Render(grid)
	}
}

func handleInsert(e ui.Event) {
	i := selected.SelectedRow
	if e.ID == "<Enter>" || e.ID == "<Escape>" {
		insertMode = false
	} else if e.ID == "<Backspace>" {
		text := selected.Rows[i]
		if len(text) > 0 {
			selected.Rows[i] = text[0 : len(text)-1]
		}
	} else if e.ID == "<Space>" {
		selected.Rows[i] += "_"
	} else {
		selected.Rows[i] += strings.ToLower(e.ID)
	}
	selectedItems[i].Name = selected.Rows[i]
}

func (gm *GymModel) normalEvents(e ui.Event) {
	switch e.ID {
	case "q", "<C-c>":
		ui.Close()
		fmt.Println("")
		fmt.Println("")
		fmt.Println("")
		fmt.Println("")
		os.Exit(1)
		return
	case "s":
		gm.Gym.RemoveStruct(gm.Name)
		gm.Gym.AddStruct(gm.Name)
		for _, item := range selectedItems {
			gm.Gym.AddFieldToStruct(gm.Name, item.Name, item.Flavor, "")
		}
		gm.Gym.Save()
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
		val := flavor.Generate(item.Flavor, "")
		rendered.Rows = append(rendered.Rows, val)
	}
}

func handleEnterFlavors() {
	if flavors.SelectedRow == 0 {
		addToSelectedItems("first_name", "first_name", "")
		addToSelectedItems("last_name", "last_name", "")
		addToSelectedItems("pronouns", "pronouns", "")
		addToSelectedItems("phone", "phone", "")
		addToSelectedItems("email", "email", "")
		addToSelectedItems("int", "age", "")
	} else if flavors.SelectedRow == 1 {
		addToSelectedItems("address", "address", "")
		addToSelectedItems("latitude", "latitude", "")
		addToSelectedItems("longitude", "longitude", "")
	} else if flavors.SelectedRow == 2 {
		addToSelectedItems("phone", "phone", "")
	} else if flavors.SelectedRow == 3 {
		addToSelectedItems("timestamp", "something_at", "")
	} else if flavors.SelectedRow == 4 {
		addToSelectedItems("few_words", "few_words", "")
	} else if flavors.SelectedRow == 5 {
		addToSelectedItems("int", "int", "")
	} else if flavors.SelectedRow == 6 {
		addToSelectedItems("float", "float", "")
	} else if flavors.SelectedRow == 7 {
		addToSelectedItems("bool", "is_on", "")
	} else if flavors.SelectedRow == 8 {
		addToSelectedItems("paragraph", "paragraph", "")
	} else if flavors.SelectedRow == 9 {
		addToSelectedItems("email", "email", "")
	} else if flavors.SelectedRow == 10 {
		addToSelectedItems("uuid", "id", "")
	} else if flavors.SelectedRow == 11 {
		addToSelectedItems("model", "created_by", "user")
	}
	//selected.SelectedRow = len(selected.Rows) + 1
}

func addToSelectedItems(flavor, name, extra string) {
	gf := GymField{}
	gf.Flavor = flavor
	gf.Name = name
	gf.Extra = extra
	selectedItems = append(selectedItems, gf)
	selected.Rows = append(selected.Rows, name)
}

func setListColors(s *widgets.List) {
	s.SelectedRowStyle.Fg = ui.ColorWhite
	s.SelectedRowStyle.Bg = ui.ColorMagenta
	s.TextStyle.Fg = ui.ColorWhite
	s.TextStyle.Bg = ui.ColorBlack
}
