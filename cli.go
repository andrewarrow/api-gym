package main

import "fmt"

func PrintHelp() {
	fmt.Println("")
	fmt.Println("  api-gym help")
	fmt.Println("  api-gym ls")
	fmt.Println("  api-gym flavors")
	fmt.Println("  api-gym model add [name]")
	fmt.Println("  api-gym model [model_index] [flavor_index_list")
	fmt.Println("  api-gym model update-random [struct_index] [field_index] [random]")
	fmt.Println("  api-gym json [model_index] [amount]")
	fmt.Println("")
	//fmt.Println("  api-gym new")
	//fmt.Println("  api-gym add      [get|post]    [route]   ")
	//fmt.Println("  api-gym rm       [route_index]           ")
	//fmt.Println("  api-gym response [route_index] [response]")
	//fmt.Println("")
	//fmt.Println("  api-gym structs update-random [struct_index] [field_index] [random]")
	//fmt.Println("")
	//fmt.Println("  api-gym run  [route_index]")
}
