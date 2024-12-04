package main

import (
	"fmt"
	"os"

	dfs "lem_in/DFS"
	A "lem_in/antJourney"
	f "lem_in/fileParsing"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println("invalid number of args ")
		return
	}
	farm := f.Readfile()
	err := f.ParseFarm(farm)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("%v\n\n%v\n\n%v\n\n%v\n\n%v\n\n", f.Fa.NOA, f.Fa.Rooms, f.Links, f.Fa.Start, f.Fa.End)
	// initialize the visited map
	for ele := range f.Fa.Rooms {
		dfs.Visited[ele] = false
	}
	// start := strings.Split(f.Fa.Start, " ")
	//	end := strings.Split(f.Farm.End, " ")
	dfs.FindPaths(&f.Fa, f.Fa.Start, "")
	for _, path := range dfs.Paths {
		fmt.Println("path : " + path)
	}
	for i := range f.Fa.NOA {
		A.AntPositions[i+1] = f.Fa.Start
	}
	fmt.Println(A.AntPositions)
}

/* func InitializeMap() {
	for _, ele := range f.Farm.Rooms {
		Visited[ele] = false
	}
} */
