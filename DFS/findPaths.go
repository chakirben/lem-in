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

// GetUniqueSortedPaths filters unique paths and sorts them by length
func GetUniqueSortedPaths(farm *f.Farm) [][]string {
	visited := InitializeMap(farm)
	allPaths := FindPaths(farm, farm.Start, []string{}, visited)

	uniquePathsMap := make(map[string]struct{})
	for _, path := range allPaths {
		// Convert path slice to a string representation
		pathStr := strings.Join(path, "-")
		uniquePathsMap[pathStr] = struct{}{}
	}

	// Convert map keys back to slices and collect unique paths
	var uniquePaths [][]string
	for pathStr := range uniquePathsMap {
		uniquePaths = append(uniquePaths, strings.Split(pathStr, "-"))
	}

	// Sort the unique paths by their length
	sort.Slice(uniquePaths, func(i, j int) bool {
		return len(uniquePaths[i]) < len(uniquePaths[j])
	})

	return uniquePaths
}
