package screen

import (
	"api-gym/files"
	"encoding/json"
	"fmt"
	"log"
	"os"

	ui "github.com/gizak/termui/v3"
	"github.com/gizak/termui/v3/widgets"
)

var lists = []*widgets.List{}

func ReadJson(file string) {
	if err := ui.Init(); err != nil {
		log.Fatalf("failed to initialize termui: %v", err)
	}
	defer ui.Close()

	lists = append(lists, widgets.NewList())
	setListColors(lists[0])

	jsonString := files.ReadFile(file)
	var m map[string]any
	json.Unmarshal([]byte(jsonString), &m)
	lists[0].Rows = []string{fmt.Sprintf("%d", len(m))}

	termWidth, termHeight := ui.TerminalDimensions()
	grid.SetRect(0, 0, termWidth, termHeight)

	grid.Set(
		ui.NewRow(1.0,
			ui.NewCol(0.16, lists[0]),
		),
	)

	ui.Render(grid)
	uiEvents := ui.PollEvents()
	for {
		e := <-uiEvents
		jsonEvents(e)
		ui.Render(grid)
	}
}

func jsonEvents(e ui.Event) {
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
		//selectedList().ScrollDown()
	case "k", "<Up>":
		//selectedList().ScrollUp()
	case "<Enter>":
		//handleEnter()
	case "<Resize>":
		payload := e.Payload.(ui.Resize)
		grid.SetRect(0, 0, payload.Width, payload.Height)
		ui.Clear()
	}
}
