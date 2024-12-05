package algo

import "lem-in/server"

// Tunnel represents a one-way tunnel in the graph.
type Tunnel struct {
	to             string
	capacity, flow int
}

// Farm represents a graph of tunnels and rooms.
var Farm = make(map[string][]Tunnel)

func AddTunnel(farm map[string][]Tunnel, from, to string, capacity int) {
	farm[from] = append(farm[from], Tunnel{to: to, capacity: capacity, flow: 0})
}

// transforms the Rooms map into a graph with capacities.
func BuildFarmFromRooms() {
	for room, connections := range server.Rooms {
		for _, connectedRoom := range connections {
			// making sure that only one ant can pass through a tunnel
			AddTunnel(Farm, room, connectedRoom, 1)
		}
	}

	// multiple ants in start and end rooms
	for _, neighbor := range server.Rooms[server.Start] {
		AddTunnel(Farm, server.Start, neighbor, len(server.Rooms)) // Multiple ants can leave start
	}

	for _, neighbor := range server.Rooms[server.End] {
		AddTunnel(Farm, neighbor, server.End, len(server.Rooms)) // Multiple ants can enter end
	}
}
