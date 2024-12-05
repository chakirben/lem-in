package algo

import "lem-in/server"


func BFS() (map[string]*Tunnel, bool) {
	queue := []string{server.Start}
	visited := make(map[string]bool)
	parent := make(map[string]*Tunnel)

	visited[server.Start] = true

	for len(queue) > 0 {
		current := queue[0]
		queue = queue[1:]

		// go through all tunnels from the current node
		for i := range Farm[current] {
			tunnel := &Farm[current][i]

			// see if there's remaining capacity and if the room is unvisited
			if !visited[tunnel.to] && (tunnel.capacity-tunnel.flow) > 0 {
				visited[tunnel.to] = true
				parent[tunnel.to] = tunnel

				// If we reached the end room, return the parent map
				if tunnel.to == server.End {
					return parent, true
				}

				queue = append(queue, tunnel.to)
			}
		}
	}

	// If no path is found, return false
	return nil, false
}
