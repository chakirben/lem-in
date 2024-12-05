package dfs

import "strings"

func TrimPaths(paths []string) [][]string {
	sl := [][]string{}
	for _, path := range paths {
		sl = append(sl, strings.Split(path, "-"))
	}
	return sl
}

func SortPath(paths [][]string) [][]string {
	for i := 0; i < len(paths)-1; i++ {
		for j := 0; j < len(paths)-1-i; j++ {
			if len(paths[j]) > len(paths[j+1]) {
				paths[j], paths[j+1] = paths[j+1], paths[j]
			}
		}
	}
	return paths
}
