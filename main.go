package main

import (
	"fmt"
	"os"

	f "lem_in/fileParsing"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println("invalid number of args ")
		return
	}
	farm := f.Readfile()
	NOA, Rooms, Links, err := f.ParseFarm(farm)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(fmt.Sprintf("%v\n\n%v\n\n%v\n\n%v\n\n%v", NOA, Rooms, Links, f.StartRoom, f.EndRoom))
	// f.CheckEroors(arr)
}
