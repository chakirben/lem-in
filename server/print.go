package server

import "fmt"

func PrintAll() {
	fmt.Println(Ants)
	fmt.Println("##start")
	fmt.Printf("%v %v %v\n", Start, Coordinates[Start][0], Coordinates[Start][1])
	for k := range Rooms {
		fmt.Printf("%v %v %v\n", k, Coordinates[k][0], Coordinates[k][1])
	}
	fmt.Println("##end")
	fmt.Printf("%v %v %v\n", End, Coordinates[End][0], Coordinates[End][1])
	for k, v := range Rooms {
		fmt.Printf("%v-%v\n", k, v)
	}
}
