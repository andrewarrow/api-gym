package main

import (
	"api-gym/gym"
	"api-gym/simulate"
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

	if command == "new" {
		g.SaveBackup()
		g = gym.NewGym()
		fmt.Printf("\nNew Gym Create with name: %s\n\n", g.Name)
		g.Save()
	} else if command == "ls" {
		g.ListRoutes()
	} else if command == "add" {
		verb := os.Args[2]
		route := os.Args[3]
		g.AddRoute(verb, route)
	} else if command == "rm" {
		index := os.Args[2]
		g.RemoveRoute(index)
	} else if command == "response" {
		index := os.Args[2]
		g.AddResponseToRoute(index, os.Args[3])
	} else if command == "run" {
		index := os.Args[2]
		simulate.Run(index, g)
	} else if command == "structs" {
		if os.Args[2] == "add" {
			name := os.Args[3]
			g.AddStruct(name)
			return
		}
		index := os.Args[2]
		name := os.Args[3]
		flavor := os.Args[4]
		random := os.Args[5]
		g.AddFieldToStruct(index, name, flavor, random)
	} else if command == "help" {
		PrintHelp()
	}
}
