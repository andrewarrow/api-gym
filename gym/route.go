package gym

type Route struct {
	Verb  string `json:"verb"`
	Route string `json:"route"`
}

func NewRoute(verb, route string) *Route {
	r := Route{}
	r.Verb = verb
	r.Route = route
	return &r
}
