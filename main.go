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
	sl := dfs.TrimPaths(dfs.Paths)
	sl2 := dfs.SortPath(sl)
	for _, path := range sl2 {
		fmt.Println(path)
	}
	A.SpreadAnts(f.Fa.NOA, sl2)
	//	A.DeleteStart()
	A.InitPathlength(sl2)
	A.InitPositions(f.Fa.Start, f.Fa.NOA)
	A.InitRooms(f.Roommss, f.Fa.Start, f.Fa.NOA)
	fmt.Println(A.AntPositions)
	fmt.Println("---------------------------------")
	fmt.Println("the Positions : ", A.AntPositions)
	fmt.Println("the path lengths : ", A.PathwithAnts)
	fmt.Println("the ant paths : ", A.AntPaths)
	fmt.Println("the rooms : ", A.RoomsSate)
	fmt.Println("---------------------------------")
	A.FindAntJourney(f.Fa.NOA, f.Fa.End, f.Fa.Start)
	for i := range f.Fa.NOA {
		A.AntPositions[i+1] = f.Fa.Start
	}
}

/* func InitializeMap() {
	for _, ele := range f.Farm.Rooms {
		Visited[ele] = false
	}
} */
