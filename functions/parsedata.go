package lem_in

import (
	"fmt"
	"strings"
)

var (
	StartRoom     string
	Started       bool = false
	EndRoom       string
	Ended         bool = false
	Rooms         []string
	Links         []string
	Coordinations []string
)

func ParseFarm(input []string) (NumberOfAnts int, rooms []string, links []string, err error) {
	DeleteComments(input)
	NOA, err2 := CheckNOfAnts(input[0])
	if err2 != nil {
		return NOA, nil, nil, err2
	}
	for i, line := range input {
		if i == 0 {
			continue
		}
		if line == "##start" || line == "##end" {
			if !CheckStartEnd(line, i, input) {
				return NOA, nil, nil, err2
			}
		}
		if CheckRoom(line) {
			if IsRoomRepeated(line) {
				return NOA, nil, nil, fmt.Errorf("Rooms Repeated")
			}
			if IsCooRepeated(line) {
				return NOA, nil, nil, fmt.Errorf("Coordinates Repeated")
			}
			ro := strings.Split(line, " ")
			Rooms = append(Rooms, ro[0])
			Coordinations = append(Coordinations, ro[1]+" "+ro[2])
			continue
		}
		if Checklink(line) {
			links = append(links, line)
			continue
		}
		return 0, nil, nil, fmt.Errorf("ERROR: invalid data format ")
	}
	return 0, nil, nil, fmt.Errorf("ERROR: invalid data format ")
}
