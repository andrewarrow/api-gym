package main

import "fmt"

func PrintHelp() {
	fmt.Println("")
	fmt.Println("  api-gym help")
	fmt.Println("  api-gym new")
	fmt.Println("  api-gym ls")
	fmt.Println("  api-gym add      [get|post] [route]   ")
	fmt.Println("  api-gym rm       [index]              ")
	fmt.Println("  api-gym response [index]    [response]")
	fmt.Println("")
	fmt.Println("  api-gym structs [index] [name] [flavor] [random]")
	fmt.Println("  api-gym structs update-random [index] [field_index] [random]")
	fmt.Println("  api-gym structs add [name]")
	fmt.Println("")
	fmt.Println("  api-gym run [index]")
	fmt.Println("")
}
