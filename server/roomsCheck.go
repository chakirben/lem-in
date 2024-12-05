package server

import (
	"fmt"
	"os"
)

func RoomChecks() {
	for room, connections := range Rooms {
		if len(connections) == 0 || (len(connections) == 1 && (connections[0] != Start || connections[0] != End)) && room != Start && room != End {
			delete(Rooms, room)
			continue
		}
		if (room == Start || room == End) && len(connections) == 0 {
			fmt.Println("ERROR: Start and End rooms must be connected to at least one other room.")
			os.Exit(2)
		} else if (room == Start || room == End) && len(connections) == 1 {
			if room == Start && connections[0] == End {
				fmt.Println("ERROR: Start and End rooms must be connected to at least one other room.")
				os.Exit(2)
			} else if room == End && connections[0] == Start {
				fmt.Println("ERROR: Start and End rooms must be connected to at least one other room.")
				os.Exit(2)
			}
		}

	}

}
