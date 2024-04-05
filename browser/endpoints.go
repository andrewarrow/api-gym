package browser

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
		return
	}
	Document.Id("r" + r.Id).AddClass("hidden")
	r.Open = false
}
