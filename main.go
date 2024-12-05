package main

import (
	"fmt"
	"lem-in/server"
	"os"
)

func main() {
	if len(os.Args) != 2 {
		PrintArgError()
		return
	}
	filepath := os.Args[1]
	err := server.ReadFile(filepath)
	if err != nil {
		fmt.Println(err)
		return
	}
	farm := server.Rooms
	for k, v := range farm {
		fmt.Println("Room: ", k, "Linked to the following: ", v)
	}
}

func PrintArgError() {
	fmt.Println("Invalid number of arguments")
	fmt.Println("Usage: go run main.go <file_path>")
}
