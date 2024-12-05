package algo

import (
	"lem-in/server"
	"math"
)

func UpdateFlow(farm map[string][]Tunnel, parent map[string]*Tunnel) int {
	pathFlow := math.MaxInt
	current := server.End

	// Find the bottleneck capacity in the path
	for current != server.Start {
		tunnel := parent[current]
		if tunnel == nil {
			break
		}
		pathFlow = min(pathFlow, tunnel.capacity-tunnel.flow)
		current = reverseFind(farm, tunnel) // Find the previous node(room)
	}

	// Update the flow(tunnel capacity) along the path
	current = server.End
	for current != server.Start {
		tunnel := parent[current]
		if tunnel == nil {
			break
		}
		tunnel.flow += pathFlow
		current = reverseFind(farm, tunnel)
	}

	return pathFlow
}

func reverseFind(farm map[string][]Tunnel, tunnel *Tunnel) string {
	for node, tunnels := range farm {
		for _, t := range tunnels {
			if t.to == tunnel.to && t.capacity == tunnel.capacity && t.flow == tunnel.flow {
				return node
			}
		}
	}
	return ""
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
