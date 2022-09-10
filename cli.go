package main

import "fmt"

func PrintHelp() {
	fmt.Println("")
	fmt.Println("  api-gym help")
	fmt.Println("  api-gym new")
	fmt.Println("  api-gym ls")
	fmt.Println("  api-gym add      [get|post]    [route]   ")
	fmt.Println("  api-gym rm       [route_index]           ")
	fmt.Println("  api-gym response [route_index] [response]")
	fmt.Println("")
	fmt.Println("  api-gym structs [struct_index] [name] [flavor] [random]")
	fmt.Println("  api-gym structs update-random [struct_index] [field_index] [random]")
	fmt.Println("  api-gym structs add [name]")
	fmt.Println("")
	fmt.Println("  api-gym run  [route_index]")
	fmt.Println("  api-gym json [struct_index] [amount]")
	fmt.Println("")
}
