package main

import (
	"fmt"
	"os"
	"strings"

	dfs "lem_in/DFS"
	f "lem_in/fileParsing"
)

func main() {
	if len(os.Args) != 2 {
		ArgError()
		return
	}
	farm := f.Readfile()
	NOA, Rooms, Links, err := f.ParseFarm(farm)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("%v\n\n%v\n\n%v\n\n%v\n\n%v\n\n", NOA, Rooms, Links, f.StartRoom, f.EndRoom)
	dfs.InitializeMap()
	start := strings.Split(f.StartRoom, " ")
	end := strings.Split(f.EndRoom, " ")
	dfs.FindPaths(Links, start[0], end[0], start[0], "")
	for _, path := range dfs.Paths {
		fmt.Println("path : " + path)
	}
}

func ArgError() {
	fmt.Println("invalid number of args")
	fmt.Println("usage: ./lem-in <filename> || go run main.go <filename>")
}
