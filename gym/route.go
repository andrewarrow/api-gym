package gym

import (
	"fmt"
)

type Route struct {
	Verb    string `json:"verb"`
	Route   string `json:"route"`
	Model   string `json:"model"`
	Count   int    `json:"count"`
	UseFile string `json:"use_file"`
}

func NewRoute(verb, route, model string, count int) *Route {
	r := Route{}
	r.Verb = verb
	r.Route = route
	r.Model = model
	r.Count = count
	r.UseFile = ""
	return &r
}

func (r *Route) String() string {
	return fmt.Sprintf("%4s %s", r.Verb, r.Route)
}

func (r *Route) ParseResponse() (string, string) {
	//structure := ""
	//flavor := ""
	//if strings.HasPrefix(r.Response, "[]") {
	//	structure = "array"
	//	flavor = r.Response[2:]
	//}

	//return structure, flavor
	return "", ""
}
