package main

import (
	"api-gym/files"
	"api-gym/flavor"
	"api-gym/gym"
	"api-gym/server"
	"api-gym/simulate"
	"api-gym/util"
	"fmt"
	"math/rand"
	"os"
	"time"
)

var g *gym.Gym

func main() {
	rand.Seed(time.Now().UnixNano())

	if len(os.Args) == 1 {
		PrintHelp()
		return
	}
	command := os.Args[1]
	g = gym.LoadGym()
	files.MkdirJson()

	if command == "new" {
		g.SaveBackup()
		g = gym.NewGym()
		fmt.Printf("\nNew Gym Create with name: %s\n\n", g.Name)
		g.Save()
	} else if command == "ls" {
		g.ListRoutes()
	} else if command == "route" {
		if os.Args[2] == "add" {
			verb := os.Args[3]
			route := os.Args[4]
			modelIndex := os.Args[5]
			g.AddRoute(verb, route, modelIndex)
			return
		}
		index := os.Args[3]
		simulate.Run(index, g)
	} else if command == "rm" {
		index := os.Args[2]
		g.RemoveRoute(index)
	} else if command == "response" {
		index := os.Args[2]
		g.AddResponseToRoute(index, os.Args[3])
	} else if command == "json" {
		structIndex := os.Args[2]
		extra := os.Args[3]
		simulate.Json(structIndex, extra, g)
	} else if command == "model" {
		if os.Args[2] == "add" {
			name := os.Args[3]
			g.AddStruct(name)
			return
		}
		if os.Args[2] == "update-random" {
			index := os.Args[3]
			fieldIndex := os.Args[4]
			g.UpdateStructRandom(index, fieldIndex, os.Args[5])
			return
		}
		modelIndex := os.Args[2]
		name := os.Args[3]
		otherModel := ""
		flavorIndex := ""
		if len(os.Args) > 4 {
			otherModel = os.Args[4]
		} else {
			flavor.ListFlavors()
			flavorIndex = util.InputLine()
		}

		g.AddFieldToStruct(modelIndex, name, otherModel, flavorIndex)
	} else if command == "flavors" {
		flavor.ListFlavors()
	} else if command == "server" {
		router := server.Setup(g)
		router.Run(":8080")
	} else if command == "help" {
		PrintHelp()
	}
}
