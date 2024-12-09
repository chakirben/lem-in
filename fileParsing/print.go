package lem_in

import "fmt"

func PrintAll() {
	fmt.Println(Fa.Ants)
	fmt.Println("##start")
	fmt.Printf("%v %v %v\n", Fa.Start, Fa.Coordinates[Fa.Start][0], Fa.Coordinates[Fa.Start][1])
	for k := range Fa.Rooms {
		if k == Fa.Start || k == Fa.End {
			continue
		}
		fmt.Printf("%v %v %v\n", k, Fa.Coordinates[k][0], Fa.Coordinates[k][1])
	}
	fmt.Println("##end")
	fmt.Printf("%v %v %v\n", Fa.End, Fa.Coordinates[Fa.End][0], Fa.Coordinates[Fa.End][1])
	for k, v := range Fa.Rooms {
		for _, link := range v {
			fmt.Printf("%v-%v\n", k, link)
		}
	}
}
