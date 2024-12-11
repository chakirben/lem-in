package main

import (
	"fmt"
	"os"

	dfs "lem_in/DFS"
	algo "lem_in/antJourney"
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
	// f.PrintAll()
	paths, filteredPaths := dfs.GetUniqueAndFilteredPaths(&f.Fa)
	for _, path := range paths {
		fmt.Println(path)
	}
	for _, path := range filteredPaths {
		fmt.Println(path)
	}

	if len(paths) == 0 {
		fmt.Println("ERROR: invalid data format, No Paths Found")
		return
	}
	// fmt.Println(paths)

	algo.Sendants(filteredPaths)
}
