package algo

import (
	"fmt"
	"lem-in/server"
)

func AntMovement(paths [][]string) {
	ants := make([]string, server.Ants)
	for i := 0; i < server.Ants; i++ {
		// // initial state of ants waiting at at the start room
		ants[i] = "waiting"
	}

	roomOccupancy := make(map[string]int) // follow room occupancy if full or not
	timeStep := 0

	for {
		timeStep++
		allFinished := true

		for i := 0; i < server.Ants; i++ {
			if ants[i] == "finished" {
				continue
			}

			allFinished = false
			antPath := paths[i%len(paths)] // Assign paths cyclically

			var currentRoomIndex int
			if ants[i] == "waiting" {
				currentRoomIndex = -1
			} else {
				for j, room := range antPath {
					if room == ants[i] {
						currentRoomIndex = j
						break
					}
				}
			}

			// Determine the next room
			var nextRoom string
			if currentRoomIndex+1 < len(antPath) {
				nextRoom = antPath[currentRoomIndex+1]
			}

			// Move if the next room is unoccupied
			if nextRoom != "" && (nextRoom == server.End || roomOccupancy[nextRoom] == 0) {
				if ants[i] != "waiting" {
					roomOccupancy[ants[i]]-- // Free current room
				}
				ants[i] = nextRoom
				roomOccupancy[nextRoom]++ // Occupy the next room

				// Print ant movement
				if nextRoom == server.Start {
					continue
				}

				fmt.Printf("L%d-%s ", i+1, nextRoom)

				if nextRoom == server.End {
					ants[i] = "finished"
				}
			}
		}

		if allFinished {
			break
		}

		fmt.Println()
	}

}
