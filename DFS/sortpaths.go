package dfs

import (
	"strings"
	"sort"
)
type pathWithScore struct {
    path  []string
    score int
}

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

func FilterPath(paths [][]string) [][]string {
	scores := []int{}
	FilteredPaths := [][]string{}
	Repititions := 0

	for i, path := range paths {
		for _, room := range path {
			Repititions += checkRep(room, Strip(i, paths))
		}
		scores = append(scores, Repititions)
		Repititions = 0
	}
/* 	for i := 0; i < len(scores)-1; i++ {
		for j := i + 1; j < len(scores); j++ {
			if scores[i] > scores[j] {
				paths[j], paths[i] = paths[i], paths[j]
			}
		}
	} */
	combined := make([]pathWithScore, len(paths))
    for i := range paths {
        combined[i] = pathWithScore{path: paths[i], score: scores[i]}
    }

    // Sort based on the score
    sort.Slice(combined, func(i, j int) bool {
        return combined[i].score < combined[j].score
    })

    // Extract sorted paths
    sortedPaths := make([][]string, len(paths))
    for i, entry := range combined {
        sortedPaths[i] = entry.path
    }
	
	FilteredPaths = append(FilteredPaths, sortedPaths[0])
	for i, path := range sortedPaths {
		if i == 0 {
			continue
		}
		if Disjoint(path, FilteredPaths) {
			FilteredPaths = append(FilteredPaths, path)
		}
	}
	return FilteredPaths
}

func Disjoint(p []string, paths [][]string) bool {
	for _, path := range paths {
		for i, room := range path {
			if i == 0 || i == len(path)-1 {
				continue
			}
			for _, r := range p {
				if room == r {
					return false
				}
			}
		}
	}
	return true
}

func checkRep(room string, paths [][]string) int {
	count := 0
	for _, path := range paths {
		for _, r := range path {
			if room == r {
				count++
			}
		}
	}
	return count
}

func Strip(i int, paths [][]string) [][]string {
	sl := [][]string{}
	for indice, path := range paths {
		if indice == i {
			continue
		}
		sl = append(sl, path)
	}
	return sl
}
