package dfs

import (
	"strings"

	P "lem_in/fileParsing"
)

var visited map[string]bool

func findShortestPath(links []string, Start string, End string, room string) bool {
	if room == End {
		return true
	}
	visited[room] = true
	neighbors := GetAdjacencyOf(room)
	for _, nei := range neighbors {
		if visited[nei] {
			continue
		}
		findShortestPath(links, Start, End, nei)
	}
}

func GetAdjacencyOf(room string) []string {
	Adjacency := []string{}
	for _, link := range P.Links {
		arr := strings.Split(link, "-")
		if room == arr[0] {
			Adjacency = append(Adjacency, arr[1])
		}
		if room == arr[1] {
			Adjacency = append(Adjacency, arr[0])
		}
	}
	return Adjacency
}
