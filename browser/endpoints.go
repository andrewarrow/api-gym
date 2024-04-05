package browser

type Route struct {
	Id string
}

func RegisterEndpoints() {
	//Document.Document.Call("addEventListener", "keydown", js.FuncOf(g.GymKeyPress))
	div := Document.ByIdWrapped("route")
	for _, input := range div.SelectAllByClass("cursor-pointer") {
		route := Route{}
		route.Id = input.Id
		input.EventWithId(route.Click)
	}
}

func (r *Route) Click() {
}
