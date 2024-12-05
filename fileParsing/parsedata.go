package lem_in

import (
	"fmt"
	"strings"
)

type Farm struct {
	Rooms map[string][]string
	Start string
	End   string
	NOA   int
}

var (
	Fa = Farm{
		Rooms: make(map[string][]string),
	}
	Rooms = make(map[string][]string)
)

var (
	Started       bool = false
	Ended         bool = false
	Links         []string
	Roommss       []string
	Coordinations []string
)

func ParseFarm(input []string) (err error) {
	input = DeleteComments(input)
	err2 := CheckNOfAnts(input[0])
	if err2 != nil {
		return err2
	}
	for i, line := range input {
		if i == 0 {
			continue
		}
		if CheckRoom(line) {
			if IsRoomRepeated(line) {
				return fmt.Errorf("Rooms Repeated")
			}
			if IsCooRepeated(line) {
				return fmt.Errorf("Coordinates Repeated")
			}
			ro := strings.Split(line, " ")
			Roommss = append(Roommss, ro[0])
			Fa.Rooms[ro[0]] = []string(nil)
			Coordinations = append(Coordinations, ro[1]+" "+ro[2])
			continue
		}

		if Checklink(line) {
			rr := strings.Split(line, "-")
			Fa.Rooms[rr[0]] = append(Fa.Rooms[rr[0]], rr[1])
			Fa.Rooms[rr[1]] = append(Fa.Rooms[rr[1]], rr[0])
			Links = append(Links, line)
			continue
		}
		if line == "##start" || line == "##end" {
			if !CheckStartEnd(line, i, input) {
				return fmt.Errorf("Error: Invalid data after ##Start or  ##End Argument")
			} else {
				continue
			}
		}
		return fmt.Errorf("ERROR: invalid data format ")
	}
	return nil
}

/* func FillAdjacency() {
	for _, ele := range Rooms {
		Rooms_Links[ele] = dfs.GetAdjacencyOf(ele)
	}
} */
