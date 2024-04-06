package app

import (
	"sort"
	"strings"

	"codeberg.org/andrewarrow/roll"
	"github.com/andrewarrow/feedback/router"
)

func HandleGym(c *router.Context, second, third string) {
	if second == "" && third == "" && c.Method == "GET" {
		handleGymIndex(c)
		return
	}
	if second == "routes" && third == "" && c.Method == "GET" {
		handleRoutes(c)
		return
	}
	if second == "json" && third == "" && c.Method == "GET" {
		handleGymJson(c)
		return
	}
	if second == "route" && third == "" && c.Method == "PUT" {
		handleRouteUpsert(c)
		return
	}
	c.NotFound = true
}

func handleGymIndex(c *router.Context) {
	send := map[string]any{}
	tops, m := getEndpoints("")
	send["items"] = tops
	send["map"] = m
	c.SendContentInLayout("endpoints.html", send, 200)
}
func getEndpoints(top string) ([]string, map[string][]map[string]any) {

	items := roll.Many(indexNameRoutes, "workspace.keyword", workspaceGuid)
	m := map[string][]map[string]any{}
	tops := []string{}
	for _, item := range items {
		route := item["route"].(string)
		tokens := strings.Split(route, "/")
		topLevel := tokens[1]
		m[topLevel] = append(m[topLevel], item)
	}
	for k, v := range m {
		tops = append(tops, k)
		sort.Slice(v, func(i, j int) bool {
			r1 := v[i]["route"].(string)
			r2 := v[j]["route"].(string)
			if r1 == r2 {
				r1 := v[i]["verb"].(string)
				r2 := v[j]["verb"].(string)
				return r1 < r2
			}
			return r1 < r2
		})
	}
	if top != "" {
		for k, _ := range m {
			if k == top {
				continue
			}
			delete(m, k)
		}
		return []string{top}, m
	}
	sort.Strings(tops)
	return tops, m
}

func handleGymJson(c *router.Context) {

	send := map[string]any{}
	c.SendContentInLayout("gym.html", send, 200)
}
