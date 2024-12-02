package server

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

var AdvMap struct {
	Ants      int
	Rooms     map[string][]string
	Cordinate map[string][]int
}

func init() {
	AdvMap.Rooms = make(map[string][]string)
	AdvMap.Cordinate = make(map[string][]int)
}

func ReadFile(filepath string) error {
	file, err := os.Open(filepath)
	if err != nil {
		return err
	}
	defer file.Close()

	firstline := true
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		// fmt.Println(line)
		if firstline {
			AdvMap.Ants, err = strconv.Atoi(line)
			if err != nil {
				return err
			}
			firstline = false
			continue
		}
		if line == "" {
			continue
		} else if line[0] == '#' {
			if line == "##start" {
				AdvMap.Rooms["start"] = []string{}
			} else if line == "##end" {
				AdvMap.Rooms["end"] = []string{}
			} else {
				continue
			}
		} else {
			if strings.Contains(line, "-") {
				// fmt.Println(line)
				// fmt.Println(strings.Split(line, "-"))
				rooms := strings.Split(line, "-")
				if _, ok := AdvMap.Rooms[rooms[0]]; !ok {
					AdvMap.Rooms[rooms[0]] = []string{}
				}
				if _, ok := AdvMap.Rooms[rooms[1]]; !ok {
					AdvMap.Rooms[rooms[1]] = []string{}
				}
				AdvMap.Rooms[rooms[0]] = append(AdvMap.Rooms[rooms[0]], rooms[1])
				AdvMap.Rooms[rooms[1]] = append(AdvMap.Rooms[rooms[1]], rooms[0])
			} else {
				room := strings.Split(line, " ")
				x, _ := strconv.Atoi(room[1])
				y, _ := strconv.Atoi(room[2])
				AdvMap.Cordinate[room[0]] = []int{x, y}
			}
		}
	}
	return nil
}
