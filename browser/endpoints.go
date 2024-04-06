package browser

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/andrewarrow/feedback/wasm"
)

var routeById map[string]*Route
var routeSelected string

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
		inputs := form.NoClearInputs()
		inputs["id"] = routeSelected
		asString, code := wasm.DoPut("/gym/route", inputs)
		var m map[string]any
		json.Unmarshal([]byte(asString), &m)
		if code == 200 {
			fmt.Println(m)
			items := m["items"].([]any)
			top := items[0].(string)
			div := Document.Id("top-" + top)
			div.Set("innerHTML", "fo")
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
	routeSelected = ""
	if r.Open == false {
		Document.Id("r" + r.Id).RemoveClass("hidden")
		r.Open = true
		Document.Id("mode").Set("innerHTML", "Edit")
		rt := Document.Id("rt" + r.Id).Get("innerHTML")
		rv := Document.Id("rv" + r.Id).Get("innerHTML")
		Document.Id("route").Set("value", strings.TrimSpace(rt))
		Document.Id("verb").Set("value", strings.TrimSpace(rv))

		routeSelected = r.Id
		return
	}
	Document.Id("r" + r.Id).AddClass("hidden")
	r.Open = false
}
