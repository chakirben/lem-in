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
	err := f.ReadFile(os.Args[1])
	if err != "" {
		fmt.Println(err)
		return
	}

	// initialize the visited map
	//visited := dfs.InitializeMap(&f.Fa)
	//dfs.FindPaths(&f.Fa, f.Fa.Start, []string{}, visited)
	paths, filterdpaths := dfs.GetUniqueAndFilteredPaths(&f.Fa)
	if len(paths) == 0 {
		fmt.Println("ERROR: invalid data format, No Paths Found")
		return
	}
	//sl := dfs.TrimPaths(dfs.Paths)
	//sl2 := dfs.SortPath(paths)
	for _, path := range paths {
		fmt.Println(path)
	}

	//A.DeleteStart()
	A.InitPathlength(filterdpaths)
	//fmt.Println("Test Test: ", A.PathwithAnts)
	//A.InitPositions(f.Fa.Start, f.Fa.Ants)
	A.SpreadAnts(f.Fa.Ants, filterdpaths)
	A.InitRooms(f.Roommss, f.Fa.Start, f.Fa.Ants)
	fmt.Println(A.AntPositions)
	fmt.Println("---------------------------------")
	fmt.Println("the Positions : ", A.AntPositions)
	fmt.Println("the path lengths : ", A.PathwithAnts)
	fmt.Println("the ant paths : ", A.AntPaths)
	fmt.Println("the rooms : ", A.RoomsSate)
	fmt.Println("---------------------------------")
	A.FindAntJourney(f.Fa.Ants, f.Fa.End, f.Fa.Start)
	for i := range f.Fa.Ants {
		A.AntPositions[i+1] = f.Fa.Start
	}
}

/* func InitializeMap() {
	for _, ele := range f.Farm.Rooms {
		Visited[ele] = false
	}
} */
