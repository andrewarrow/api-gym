package simulate

import (
	"api-gym/gym"
	"fmt"
	"strconv"
	"strings"
)

func Run(routeIndex string, g *gym.Gym) {
	//routeIndexAsInt, _ := strconv.Atoi(routeIndex)
	//route := g.Routes[routeIndexAsInt-1]

	//modelIndexAsInt, _ := strconv.Atoi(route.ModelIndex)
	//s := g.Structs[modelIndexAsInt-1]
	//PrintItemsToStdout(s, g)
}

func makeSingleModel(field *gym.Field, g *gym.Gym) string {
	s := g.StructsByName[field.FlavorToStructName()]

	return makeStructJson(s, g, false)
}
func makeArrayItems(field *gym.Field, g *gym.Gym) string {
	s := g.StructsByName[field.FlavorToStructName()]

	amount, _ := strconv.Atoi("1")

	sub := []string{}
	for i := 0; i < amount; i++ {
		sub = append(sub, makeStructJson(s, g, false))
	}
	return strings.Join(sub, ",")
}

func makeMapItems(field *gym.Field, g *gym.Gym) string {
	s := g.StructsByName[field.FlavorToStructName()]

	amount, _ := strconv.Atoi("1")

	sub := []string{}
	for i := 0; i < amount; i++ {
		sub = append(sub, makeStructJson(s, g, false))
	}
	return strings.Join(sub, ",")
}

func PrintItemsToString(s *gym.Struct, g *gym.Gym, count int) string {
	buff := []string{}
	buff = append(buff, fmt.Sprintf(`{"%s":`, s.JsonContainerName()))

	if s.ArrayOrMap == "array" {

		buff = append(buff, "[")
		sub := []string{}
		for i := 0; i < count; i++ {
			sub = append(sub, makeStructJson(s, g, true))
		}
		buff = append(buff, strings.Join(sub, ","))
		buff = append(buff, "]")
		buff = append(buff, "}")
	} else if s.ArrayOrMap == "map" {
		/*
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
		*/
	}

	return strings.Join(buff, "")
}

func PrintItemsToStdout(s *gym.Struct, g *gym.Gym) {
	fmt.Println(PrintItemsToString(s, g, 1))
}

func makeStructJson(s *gym.Struct, g *gym.Gym, top bool) string {
	thing := []string{}
	items := []string{}
	thing = append(thing, "{")
	for _, f := range s.Fields {
		items = append(items, makeJsonBasedOnFlavor(f, g, top))
	}
	thing = append(thing, strings.Join(items, ","))
	thing = append(thing, "}")
	return strings.Join(thing, "")
}

func makeJsonBasedOnFlavor(f *gym.Field, g *gym.Gym, top bool) string {
	dt := f.DataType()
	if dt == "string" {
		value := f.ToFakeValue()
		if value == "null" {
			return fmt.Sprintf(`"%s": null`, f.NameToJson())
		} else {
			return fmt.Sprintf(`"%s": "%s"`, f.NameToJson(), value)
		}
	} else if dt == "model" && top {
		return fmt.Sprintf(`"%s": %s`, f.NameToJson(), makeSingleModel(f, g))
	} else if dt == "[]model" {
		return fmt.Sprintf(`"%s": [%s]`, f.NameToJson(), makeArrayItems(f, g))
	} else if dt == "int" || dt == "int64" || dt == "float64" || dt == "bool" {
		return fmt.Sprintf(`"%s": %s`, f.NameToJson(), f.ToFakeValue())
	}
	//return fmt.Sprintf(`"%s": %s`, f.NameToJson(), makeMapItems(f, g))
	return fmt.Sprintf(`"%s": %s`, f.NameToJson(), "null")
}
