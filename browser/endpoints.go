package browser

import (
	"encoding/json"
	"strings"

	"github.com/andrewarrow/feedback/wasm"
)

var routeById map[string]*Route

type Route struct {
	Id   string
	Open bool
}

func RegisterEndpoints() {
	routeById = map[string]*Route{}
	//Document.Document.Call("addEventListener", "keydown", js.FuncOf(g.GymKeyPress))
	div := Document.ByIdWrapped("routes")
	for _, input := range div.SelectAllByClass("cursor-pointer") {
		route := Route{}
		route.Id = input.Id[2:]
		routeById[route.Id] = &route
		input.EventWithId(route.Click)
	}

	Global.SubmitEvent("verb-form", UpsertRoute)
}

func UpsertRoute() {
	go func() {
		form := Document.Id("verb-form")
		asString, code := wasm.DoPut("/gym/route", form.NoClearInputs())
		var m map[string]any
		json.Unmarshal([]byte(asString), &m)
		if code == 200 {
			return
		}
		flashThree(asString)
	}()
}

func (r *Route) Click() {
	Document.Id("mode").Set("innerHTML", "Add")

	div := Document.ByIdWrapped("routes")
	for _, input := range div.SelectAllByClass("cursor-pointer") {
		id := input.Id[2:]
		if r.Id == id {
			continue
		}
		Document.Id("r" + id).AddClass("hidden")
		routeById[id].Open = false
	}
	if r.Open == false {
		Document.Id("r" + r.Id).RemoveClass("hidden")
		r.Open = true
		Document.Id("mode").Set("innerHTML", "Edit")
		rt := Document.Id("rt" + r.Id).Get("innerHTML")
		Document.Id("route").Set("value", strings.TrimSpace(rt))
		return
	}
	Document.Id("r" + r.Id).AddClass("hidden")
	r.Open = false
}
