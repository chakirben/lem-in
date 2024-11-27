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
	arr := f.Readfile()
	f.CheckEroors(arr)
}
