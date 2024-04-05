package browser

type Route struct {
	Id   string
	Open bool
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
	if r.Open == false {
		Document.Id("r" + r.Id).RemoveClass("hidden")
		r.Open = true
		return
	}
	Document.Id("r" + r.Id).AddClass("hidden")
	r.Open = false
}
