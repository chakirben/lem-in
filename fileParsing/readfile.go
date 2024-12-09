package lem_in

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

type Farm struct {
	Ants        int
	Rooms       map[string][]string
	Coordinates map[string][2]int
	Start       string
	End         string
	IsStart     bool
	IsEnd       bool
}

var Roommss []string
var Fa = Farm{
	Rooms:       make(map[string][]string),
	Coordinates: make(map[string][2]int),
}

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
			return "ERROR: invalid data format"
		}

		if line[0] == '#' { // Handle comments and special start/end markers.
			switch line {
			case "##start":
				Fa.IsStart = true
				continue
			case "##end":
				Fa.IsEnd = true
				continue
			default:
				continue
			}
		}

		if line[0] == 'L' {
			return "ERROR: invalid data format, room name cannot start with 'L'"
		}

		// Get number of ants.
		if firstLine {
			Fa.Ants, err = strconv.Atoi(line)
			if err != nil || Fa.Ants <= 0 {
				return "ERROR: invalid data format, invalid number of ants"
			}
			firstLine = false
			continue
		}

		// Handle start and end rooms.
		if Fa.IsStart {
			if !SetRoom(line, true) {
				return "ERROR: invalid data format, invalid start room"
			}
			Fa.IsStart = false
			continue
		}
		if Fa.IsEnd {
			if !SetRoom(line, false) {
				return "ERROR: invalid data format, invalid end room"
			}
			Fa.IsEnd = false
			continue
		}

		// Handle rooms and tunnels.
		if strings.Contains(line, "-") {
			if !Tunnels(line) {
				return "ERROR: invalid data format, invalid tunnel"
			}
		} else {
			if !SetRoom(line, false) {
				return "ERROR: invalid data format, invalid room"
			}
		}
	}

	// Ensure valid start and end rooms are set.
	if Fa.Start == "" || Fa.End == "" || Fa.Start == Fa.End {
		return "ERROR: invalid data format, Start or End room not found, or share the same room"
	}

	return ""
}

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

	Fa.Coordinates[roomName] = [2]int{x, y}
	if _, exists := Fa.Rooms[roomName]; !exists {
		Roommss = append(Roommss, roomName)
		Fa.Rooms[roomName] = []string{}
	}

	if isStart {
		Roommss = append(Roommss, roomName)
		Fa.Start = roomName
	} else if Fa.IsEnd {
		Roommss = append(Roommss, roomName)
		Fa.End = roomName
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
	if _, exists := Fa.Rooms[room1]; !exists {
		Fa.Rooms[room1] = []string{}
	}
	if _, exists := Fa.Rooms[room2]; !exists {
		Fa.Rooms[room2] = []string{}
	}

	// Avoid duplicate connections.
	if !contains(Fa.Rooms[room1], room2) {
		Fa.Rooms[room1] = append(Fa.Rooms[room1], room2)
	}
	if !contains(Fa.Rooms[room2], room1) {
		Fa.Rooms[room2] = append(Fa.Rooms[room2], room1)
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
