package simulate

import (
	"api-gym/gym"
	"fmt"
	"strconv"
	"strings"
)

func Run(index string, g *gym.Gym) {
	indexAsInt, _ := strconv.Atoi(index)
	route := g.Routes[indexAsInt]
	structure, flavor := route.ParseResponse()

	if structure == "array" {
		s := g.StructsByName[flavor]
		printItems(s, "9", g)
		return
	}
}

func makeArrayItems(field *gym.Field, g *gym.Gym) string {
	s := g.StructsByName[field.FlavorToStructName()]

	amount, _ := strconv.Atoi(field.Extra)

	sub := []string{}
	for i := 0; i < amount; i++ {
		sub = append(sub, makeStructJson(s, g))
	}
	return strings.Join(sub, ",")
}

func makeMapItems(field *gym.Field, g *gym.Gym) string {
	s := g.StructsByName[field.FlavorToStructName()]

	amount, _ := strconv.Atoi(field.Extra)

	sub := []string{}
	for i := 0; i < amount; i++ {
		sub = append(sub, makeStructJson(s, g))
	}
	return strings.Join(sub, ",")
}

func printItems(s *gym.Struct, extra string, g *gym.Gym) {

	buff := []string{}
	buff = append(buff, fmt.Sprintf(`{"%s":`, s.JsonContainerName()))

	if s.ArrayOrMap == "array" {

		buff = append(buff, "[")
		sub := []string{}
		amountAsInt, _ := strconv.Atoi(extra)
		for i := 0; i < amountAsInt; i++ {
			sub = append(sub, makeStructJson(s, g))
		}
		buff = append(buff, strings.Join(sub, ","))
		buff = append(buff, "]")
		buff = append(buff, "}")
	} else if s.ArrayOrMap == "map" {
		buff = append(buff, "{")
		sub := []string{}

		extraTokens := strings.Split(extra, ",")
		for _, token := range extraTokens {
			subFields := []string{}
			for _, f := range s.Fields {
				subFields = append(subFields, makeJsonBasedOnFlavor(f, g))
			}
			sub = append(sub, fmt.Sprintf(`"%s": {%s}`, token, strings.Join(subFields, ",")))
		}
		buff = append(buff, strings.Join(sub, ","))
		buff = append(buff, "}")
		buff = append(buff, "}")
	}

	fmt.Println(strings.Join(buff, ""))
}

func makeStructJson(s *gym.Struct, g *gym.Gym) string {
	thing := []string{}
	items := []string{}
	thing = append(thing, "{")
	for _, f := range s.Fields {
		items = append(items, makeJsonBasedOnFlavor(f, g))
	}
	thing = append(thing, strings.Join(items, ","))
	thing = append(thing, "}")
	return strings.Join(thing, "")
}

func makeJsonBasedOnFlavor(f *gym.Field, g *gym.Gym) string {
	if f.Flavor == "string" {
		return fmt.Sprintf(`"%s": "%s"`, f.NameToJson(), f.ToFakeValue())
	} else if strings.HasPrefix(f.Flavor, "[]") {
		return fmt.Sprintf(`"%s": [%s]`, f.NameToJson(), makeArrayItems(f, g))
	} else if f.Flavor == "int" || f.Flavor == "int64" || f.Flavor == "float64" {
		return fmt.Sprintf(`"%s": %s`, f.NameToJson(), f.ToFakeValue())
	}
	return fmt.Sprintf(`"%s": %s`, f.NameToJson(), makeMapItems(f, g))
}
