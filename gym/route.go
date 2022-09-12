package gym

import (
	"fmt"
	"strings"
)

type Route struct {
	Verb       string `json:"verb"`
	Route      string `json:"route"`
	Response   string `json:"response"`
	ModelIndex string `json:"model_index"`
}

func NewRoute(verb, route, modelIndex string) *Route {
	r := Route{}
	r.Verb = verb
	r.Route = route
	r.ModelIndex = modelIndex
	return &r
}

func (r *Route) String() string {
	return fmt.Sprintf("%4s %s", r.Verb, r.Route)
}

func (r *Route) ParseResponse() (string, string) {
	structure := ""
	flavor := ""
	if strings.HasPrefix(r.Response, "[]") {
		structure = "array"
		flavor = r.Response[2:]
	}

	return structure, flavor
}
