package server

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

var (
	Ants      int
	Rooms     map[string][]string
	Cordinate map[string][]int
	Start    string
	End      string
)

func init() {
	Rooms = make(map[string][]string)
	Cordinate = make(map[string][]int)
}

func ReadFile(filepath string) error {
	file, err := os.Open(filepath)
	if err != nil {
		return err
	}
	defer file.Close()

	firstline := true
	startRoom := false
	endRoom := false
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		// fmt.Println(line)
		if firstline {
			Ants, err = strconv.Atoi(line)
			if err != nil {
				return err
			}
			firstline = false
			continue
		} else if startRoom {
			Start = line
			startRoom = false
		} else if endRoom {
			End = line
			endRoom = false
		} else

		if line == "" {
			continue
		} else if line[0] == '#' {
			if line == "##start" {
				startRoom = true
			} else if line == "##end" {
				endRoom = true
			} else {
				continue
			}
		} else {
			if strings.Contains(line, "-") {
				// fmt.Println(line)
				// fmt.Println(strings.Split(line, "-"))
				rooms := strings.Split(line, "-")
				if _, ok := Rooms[rooms[0]]; !ok {
					Rooms[rooms[0]] = []string{}
				}
				if _, ok := Rooms[rooms[1]]; !ok {
					Rooms[rooms[1]] = []string{}
				}
				Rooms[rooms[0]] = append(Rooms[rooms[0]], rooms[1])
				Rooms[rooms[1]] = append(Rooms[rooms[1]], rooms[0])
			} else {
				room := strings.Split(line, " ")
				x, _ := strconv.Atoi(room[1])
				y, _ := strconv.Atoi(room[2])
				Cordinate[room[0]] = []int{x, y}
			}
		}
	}
	return nil
}
