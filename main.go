package main

import (
	"api-gym/gym"
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

	if command == "new" {
		g = gym.NewGym()
		fmt.Printf("\nNew Gym Create with name: %s\n\n", g.Name)
	} else if command == "ls" {
		g.ListRoutes()
	} else if command == "add" {
		verb := os.Args[2]
		route := os.Args[3]
		g.AddRoute(verb, route)
	} else if command == "help" {
		PrintHelp()
	}
}
