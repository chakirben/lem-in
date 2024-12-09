package dfs

import (
	lem_in "lem_in/fileParsing"
	"strings"
)

// import (
// 	f "lem_in/fileParsing"
// 	"strings"
// )

func InitializeMap(fr *lem_in.Farm) map[string]bool {
	visited := make(map[string]bool)
	for room := range fr.Rooms {
		visited[room] = false
	}
	return visited
}

// FindPaths recursively finds all paths from start to end
func FindPaths(farm *lem_in.Farm, room string, path []string, visited map[string]bool) [][]string {
	if room == farm.End {
		return [][]string{append(path, farm.End)}
	}

	visited[room] = true
	var allPaths [][]string

	for _, neighbor := range farm.Rooms[room] {
		if !visited[neighbor] {
			newPaths := FindPaths(farm, neighbor, append(path, room), visited)
			allPaths = append(allPaths, newPaths...)
		}
	}
	visited[room] = false // Backtrack
	return allPaths
}

func GetUniqueAndFilteredPaths(farm *lem_in.Farm) ([][]string, [][]string) {
	visited := InitializeMap(farm)
	allPaths := FindPaths(farm, farm.Start, []string{}, visited)

	uniquePathsMap := make(map[string]struct{})
	for _, path := range allPaths {
		pathStr := strings.Join(path, "-")
		uniquePathsMap[pathStr] = struct{}{}
	}

	var uniquePaths [][]string
	for pathStr := range uniquePathsMap {
		uniquePaths = append(uniquePaths, strings.Split(pathStr, "-"))
	}

	// // Sort the unique paths by their length
	// sort.Slice(uniquePaths, func(i, j int) bool {
	// 	return len(uniquePaths[i]) < len(uniquePaths[j])
	// })

	usedRooms := make(map[string]bool)
	var FilteredPaths [][]string

	for _, path := range uniquePaths {
		isNonOverlapping := true
		for _, room := range path {
			if room != farm.Start && room != farm.End && usedRooms[room] {
				isNonOverlapping = false
				break
			}
		}

		if isNonOverlapping {
			FilteredPaths = append(FilteredPaths, path)
			// Mark rooms in the current path as used
			for _, room := range path {
				if room != farm.Start && room != farm.End {
					usedRooms[room] = true
				}
			}
		}
	}

	return uniquePaths, FilteredPaths
}

// ########################################################################## test Other way

// func GetUniqueAndFilteredPaths(farm *lem_in.Farm) ([][]string, [][][]string) {
// 	visited := InitializeMap(farm)
// 	allPaths := FindPaths(farm, farm.Start, []string{}, visited)

// 	uniquePathsMap := make(map[string]struct{})
// 	for _, path := range allPaths {
// 		pathStr := strings.Join(path, "-")
// 		uniquePathsMap[pathStr] = struct{}{}
// 	}

// 	var uniquePaths [][]string
// 	for pathStr := range uniquePathsMap {
// 		uniquePaths = append(uniquePaths, strings.Split(pathStr, "-"))
// 	}

// 	// Recursive helper function to generate all combinations of non-overlapping paths
// 	var findCombinations func(paths [][]string, usedRooms map[string]bool, currentCombination [][]string) [][][]string
// 	findCombinations = func(paths [][]string, usedRooms map[string]bool, currentCombination [][]string) [][][]string {
// 		if len(paths) == 0 {
// 			return [][][]string{currentCombination}
// 		}

// 		var results [][][]string
// 		for i, path := range paths {
// 			isNonOverlapping := true
// 			for _, room := range path {
// 				if room != farm.Start && room != farm.End && usedRooms[room] {
// 					isNonOverlapping = false
// 					break
// 				}
// 			}

// 			if isNonOverlapping {
// 				// Mark rooms as used and add the current path to the combination
// 				newUsedRooms := copyMap(usedRooms)
// 				for _, room := range path {
// 					if room != farm.Start && room != farm.End {
// 						newUsedRooms[room] = true
// 					}
// 				}
// 				newCombination := append(currentCombination, path)
// 				// Explore further combinations excluding the current path
// 				results = append(results, findCombinations(paths[i+1:], newUsedRooms, newCombination)...)
// 			}
// 		}
// 		// Add the case where no more paths are selected
// 		results = append(results, currentCombination)
// 		return results
// 	}

// 	// Initialize with an empty combination and no used rooms
// 	allCombinations := findCombinations(uniquePaths, make(map[string]bool), nil)

// 	return uniquePaths, allCombinations
// }

// // Helper function to create a copy of a map
// func copyMap(original map[string]bool) map[string]bool {
// 	copy := make(map[string]bool)
// 	for key, value := range original {
// 		copy[key] = value
// 	}
// 	return copy
// }
