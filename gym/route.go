package gym

import "fmt"

type Route struct {
	Verb     string `json:"verb"`
	Route    string `json:"route"`
	Response string `json:"response"`
}

func NewRoute(verb, route string) *Route {
	r := Route{}
	r.Verb = verb
	r.Route = route
	return &r
}

func (r *Route) String() string {
	return fmt.Sprintf("%4s %s", r.Verb, r.Route)
}
