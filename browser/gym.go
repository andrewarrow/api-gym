package browser

import (
	"encoding/json"
	"fmt"
	"strings"
	"syscall/js"

	"github.com/brianvoe/gofakeit"
)

var root map[string]any

type Gym struct {
	lineIndex int
}

var g = Gym{}

func RegisterGymEvents() {
	site := map[string]any{"name": gofakeit.HackerPhrase()}
	sites := []map[string]any{site}
	root = map[string]any{"sites": sites}
	Document.Document.Call("addEventListener", "keydown", js.FuncOf(g.GymKeyPress))
	g.GymRender()
}

func (g *Gym) findLocation() map[string]any {
	t := root["sites"].([]map[string]any)
	fmt.Println(t, fmt.Sprintf("%T", t))
	return nil
	//return t[0].(map[string]any)
}

func findPlace(at int) {
	b, _ := json.MarshalIndent(root, "", "  ")
	s := string(b)
	lines := strings.Split(s, "\n")
	// { [{} ] }
	// { { { [ ] } } }
	for i, line := range lines {
		fmt.Println(line)
		if i == at {
			break
		}
	}
}

func (g *Gym) GymRender() {
	gDoc := Document.Id("gym")

	b, _ := json.MarshalIndent(root, "", "&nbsp;&nbsp;")
	s := string(b)
	lines := strings.Split(s, "\n")
	buffer := []string{}
	for i, line := range lines {
		class := ""
		if i == g.lineIndex {
			class = "bg-blue-600"
		}
		buffer = append(buffer, fmt.Sprintf(`<div class="%s">%s</div>`, class, line))
	}
	findPlace(g.lineIndex)
	buffer = append(buffer, fmt.Sprintf(`<div class="%s">%d</div>`, "bg-pink-600",
		g.lineIndex))
	gDoc.Set("innerHTML", strings.Join(buffer, "\n"))
}

func (g *Gym) GymKeyPress(this js.Value, p []js.Value) any {
	//p[0].Call("preventDefault")
	k := p[0].Get("key").String()
	//fmt.Println(k, vim.FullScreenMode)
	if k == "Meta" || k == "Shift" || k == "Control" {
		return nil
	}
	if k == "Escape" {
	}
	if k == "Enter" {
		prompt := Global.Global.Get("prompt").Invoke("enter command", "hello|string")
		p := prompt.String()
		tokens := strings.Split(p, "|")
		if len(tokens) != 2 {
			return nil
		}
		field := tokens[0]
		flavor := tokens[1]
		m := g.findLocation()
		if flavor == "string" {
			m[field] = gofakeit.HackerPhrase()
		}

		return nil
	}

	if k == "ArrowUp" {
		g.lineIndex--
	} else if k == "ArrowDown" {
		g.lineIndex++
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
	} else if k == "0" {
	} else if k == "$" {
	} else if k == "p" {
	}

	g.GymRender()

	return nil
}
