package main

import (
	"fmt"
	"os"

	f "lem_in/functions"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println("invalid number of args ")
		return
	}
	farm := f.Readfile()
	f.ParseFarm(farm)
	//f.CheckEroors(arr)
}
