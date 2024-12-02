package dfs

import (
	"strings"

	P "lem_in/fileParsing"
)

var (
	Visited = make(map[string]bool)
	Paths   []string
)

func InitializeMap() {
	for _, ele := range P.Rooms {
		Visited[ele] = false
	}
}

func FindPaths(links []string, Start string, End string, room string, path string) {
	if room == End {
		Paths = append(Paths, path[1:]+"-"+End)
		return
	}
	Visited[room] = true
	neighbors := GetAdjacencyOf(room)
	for _, nei := range neighbors {
		if Visited[nei] {
			continue
		}
		FindPaths(links, Start, End, nei, path+"-"+room)
	}
	Visited[room] = false
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

/* func findShortestPath(links []string, Start string, End string, room string) bool {
	if room == End {
		return true
	}
	Visited[room] = true
	neighbors := GetAdjacencyOf(room)
	for _, nei := range neighbors {
		if Visited[nei] {
			continue
		}
		findShortestPath(links, Start, End, nei)
	}
	return true
}
*/
/* func InCircles(room string, path string) bool {
	pathRooms := strings.Split(path, "-")
	for _, ele := range pathRooms {
		if ele == room {
			return true
		}
	}
	return false
} */
