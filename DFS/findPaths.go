package dfs
import f"lem_in/fileParsing"

var (
	Visited = make(map[string]bool)
	Paths   []string
)

func InitializeMap(fr *f.Farm) {
	for ele := range fr.Rooms {
		Visited[ele] = false
	}
}

func FindPaths(Frm *f.Farm, room string, path string) {
	if room == Frm.End {
		Paths = append(Paths, path[1:]+"-"+Frm.End)
		return
	}
	Visited[room] = true
	// neighbors := GetAdjacencyOf(room)
	for _, nei := range Frm.Rooms[room] {
		if Visited[nei] {
			continue
		}
		FindPaths(Frm, nei, path+"-"+room)
	}
	Visited[room] = false
}

/* func GetAdjacencyOf(room string) []string {
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
} */

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
