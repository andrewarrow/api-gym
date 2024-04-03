package browser

import (
	"fmt"
	"strings"
	"syscall/js"
)

var root map[string]any = map[string]any{}

func RegisterGymEvents() {
	Document.Document.Call("addEventListener", "keydown", js.FuncOf(GymKeyPress))
}

func GymRender() {
	g := Document.Id("gym")
	buffer := []string{}
	for k, v := range root {
		buffer = append(buffer, fmt.Sprintf(`"%s": "%v"`, k, v))
	}
	s := strings.Join(buffer, ",<br/>")
	g.Set("innerHTML", "<pre>{<br>/>"+s+"<br/>}</pre>")
}

func GymKeyPress(this js.Value, p []js.Value) any {
	//p[0].Call("preventDefault")
	k := p[0].Get("key").String()
	//fmt.Println(k, vim.FullScreenMode)
	if k == "Meta" || k == "Shift" || k == "Control" {
		return nil
	}
	if k == "Escape" {
	}

	if k == "ArrowUp" {
		root["hi2"] = "Test"
	} else if k == "ArrowDown" {
		root["hi"] = "Test"
	} else if k == "ArrowRight" {
	} else if k == "ArrowLeft" {
	} else if k == "i" {
	} else if k == "r" {
	} else if k == "O" {
	} else if k == "o" {
	} else if k == "x" {
	} else if k == "V" {
	} else if k == ":" {
	} else if k == "a" {
	} else if k == "d" {
	} else if k == "D" {
	} else if k == "m" {
	} else if k == "u" {
	} else if k == "Enter" {
	} else if k == "0" {
	} else if k == "$" {
	} else if k == "p" {
	}

	GymRender()

	return nil
}
