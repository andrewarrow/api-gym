package gym

type Route struct {
	Verb  string
	Route string
}

func NewRoute(verb, route string) *Route {
	r := Route{}
	r.Verb = verb
	r.Route = route
	return &r
}
