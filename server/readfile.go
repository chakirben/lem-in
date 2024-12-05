package server

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var (
	Ants        int
	Rooms       = make(map[string][]string)
	Coordinates = make(map[string][2]int)
	Start       string
	End         string
	IsStart     bool = false
	IsEnd       bool = false
)

// ReadFile parses the input file to extract the number of ants, rooms, and connections.
func ReadFile(filepath string) string {
	file, err := os.Open(filepath)
	if err != nil {
		return "ERROR: could not open file"
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	firstLine := true

	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if line == "" {
			continue
		}

		if line[0] == '#' {
			switch line {
			case "##start":
				IsStart = true
				continue
			case "##end":
				IsEnd = true
				continue
			default:
				continue
			}
		}

		// get number of ants
		if firstLine {
			Ants, err = strconv.Atoi(line)
			if err != nil || Ants <= 0 {
				return "ERROR: invalid data format, invalid number of ants"

			}
			firstLine = false
			continue
		}

		// treating start and end rooms
		if IsStart {
			if !SetRoom(line, true) {
				return "ERROR: invalid data format, invalid start room"
			}
			IsStart = false
			continue
		}
		if IsEnd {
			if !SetRoom(line, false) {
				return "ERROR: invalid data format, invalid end room"
			}
			IsEnd = false
			continue
		}

		// Parse rooms and connections
		if strings.Contains(line, "-") {
			if !Tunnels(line) {
				return "ERROR: invalid data format, invalid Tunnel"
			}
		} else {
			if !SetRoom(line, false) {
				return "ERROR: invalid data format, invalid Room"
			}
		}

		// if strings.Contains(line, " ") {
		// 	if !SetCordinates(line) {
		// 		return "ERROR: invalid data format, invalid Cordinates"
		// 	}
		// } else {
		// 	return "ERROR: invalid data format"
		// }

		if Start == "" || End == "" {
			fmt.Println("Start: ", Start)
			fmt.Println("End: ", End)
			return "ERROR: invalid data format, Start or End room not found"
		}
	}

	return ""
}

// func SetCordinates(line string) bool {
// 	parts := strings.Fields(line)
// 	if len(parts) != 3 {
// 		return false
// 	}

// 	x, err1 := strconv.Atoi(parts[1])
// 	y, err2 := strconv.Atoi(parts[2])
// 	if err1 != nil || err2 != nil {
// 		return false
// 	}
// 	roomName := parts[0]
// 	Coordinates[roomName] = [2]int{x, y}
// 	return true
// }

// SetRoom parses a room line and optionally sets it as the start or end room.
func SetRoom(line string, isStart bool) bool {
	parts := strings.Fields(line)
	if len(parts) != 3 {
		return false
	}

	x, err1 := strconv.Atoi(parts[1])
	y, err2 := strconv.Atoi(parts[2])
	if err1 != nil || err2 != nil {
		return false
	}

	roomName := parts[0]
	Coordinates[roomName] = [2]int{x, y}
	if _, exists := Rooms[roomName]; !exists {
		Rooms[roomName] = []string{}
	}

	if isStart {
		Start = roomName
	} else if IsEnd {
		End = roomName
	}

	return true
}

// Tunnels parses a connection line and adds it to the adjacency list.
func Tunnels(line string) bool {
	rooms := strings.Split(line, "-")
	if len(rooms) != 2 {
		return false
	}

	room1, room2 := rooms[0], rooms[1]
	if _, exists := Rooms[room1]; !exists {
		Rooms[room1] = []string{}
	}
	if _, exists := Rooms[room2]; !exists {
		Rooms[room2] = []string{}
	}

	// Avoid duplicate connections
	if !contains(Rooms[room1], room2) {
		Rooms[room1] = append(Rooms[room1], room2)
	}
	if !contains(Rooms[room2], room1) {
		Rooms[room2] = append(Rooms[room2], room1)
	}

	return true
}

// contains checks if a slice contains a specific element.
func contains(slice []string, item string) bool {
	for _, v := range slice {
		if v == item {
			return true
		}
	}
	return false
}
