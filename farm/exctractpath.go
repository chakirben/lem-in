package algo

import "lem-in/server"

func ExtractPaths(farm map[string][]Tunnel) [][]string {
	var paths [][]string

	for {
		path := []string{}
		visited := make(map[string]bool)
		current := server.Start

		for current != server.End {
			path = append(path, current)
			visited[current] = true
			foundNext := false

			for i := range farm[current] {
				tunnel := &farm[current][i]
				if tunnel.flow > 0 && !visited[tunnel.to] { // Ensure no revisits
					tunnel.flow-- // Reduce flow(free tunnel)
					current = tunnel.to
					foundNext = true
					break
				}
			}

			if !foundNext {
				break
			}
		}

		// Validate and append the path
		if len(path) > 0 && current == server.End {
			path = append(path, server.End)
			paths = append(paths, path)
		} else {
			break
		}
	}

	return paths
}
