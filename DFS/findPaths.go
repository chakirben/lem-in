package dfs

import (
	f "lem_in/fileParsing"
	"sort"
	"strings"
)

func InitializeMap(fr *f.Farm) map[string]bool {
	visited := make(map[string]bool)
	for room := range fr.Rooms {
		visited[room] = false
	}
	return visited
}

// FindPaths recursively finds all paths from start to end
func FindPaths(farm *f.Farm, room string, path []string, visited map[string]bool) [][]string {
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

func GetUniqueAndFilteredPaths(farm *f.Farm) ([][]string, [][]string) {
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

	// Sort the unique paths by their length
	sort.Slice(uniquePaths, func(i, j int) bool {
		return len(uniquePaths[i]) < len(uniquePaths[j])
	})

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

