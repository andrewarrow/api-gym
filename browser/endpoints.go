package browser

type Route struct {
	Id string
}

func RegisterEndpoints() {
	//Document.Document.Call("addEventListener", "keydown", js.FuncOf(g.GymKeyPress))
	div := Document.ByIdWrapped("routes")
	for _, input := range div.SelectAllByClass("cursor-pointer") {
		route := Route{}
		route.Id = input.Id[2:]
		input.EventWithId(route.Click)
	}
}

func (r *Route) Click() {
	Document.Id("r" + r.Id).RemoveClass("hidden")
}
