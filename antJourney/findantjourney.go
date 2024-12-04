package ants

import "fmt"

var (
	AntPositions = make(map[int]string)
	AntPaths     = make(map[int][]string)
	RoomsSate    = make(map[string]int)
)

func FindAntJourney(NOA string, paths []string, end string) {
	for i := range NOA {
		if Move(i+1, end) {
			fmt.Printf("L%v-%v ", i, AntPositions[i])
		}
	}
	fmt.Println()
}

func Move(i int, end string) bool {
	if len(AntPaths[i]) == 0 {
		return false
	}
	if AntPaths[i][0] == end {
		AntPositions[i] = AntPaths[i][0]
		AntPaths[i] = nil
		RoomsSate[AntPaths[i][0]]++
		return true
	}
	if RoomsSate[AntPaths[i][0]] == 0 {
		AntPositions[i] = AntPaths[i][0]
		AntPaths[i] = AntPaths[i][1:]
		RoomsSate[AntPaths[i][0]] = 1
		return true
	}
	return false
}

func SpreadAnts(NOA string, paths [][]string) {
	PathIndice := 0
	CurrentPath := paths[PathIndice]
	Pathlength := len(CurrentPath)
	for i := range NOA {
		if i == 0 {
			AntPaths[i+1] = CurrentPath
			Pathlength++
		}
		if Pathlength > len(paths[PathIndice+1]) {
			CurrentPath = paths[PathIndice+1]
			PathIndice++
		}
		AntPaths[i+1] = CurrentPath
	}
}

func IsValidStep() {
}

/*how to do this
in my mind i well do somthing lik range for all path and check is valid step for the ants
if this is good idea  i dont now but i work on it if i have a problem in my logic tell me
*/

/*explain the logic */

/*
the function check step by step if the step is valid the ant go in this roms but if not valid the ant break one step
*/
/* 
 func TakeYourPath(ant int) {
	for i, room := range AntPaths[ant] {
		if i == 0 {
			continue
		}
		if RoomsSate[room] == 0 {
			RoomsSate[room] = ant
			room
			AntPositions[ant] = room

		}
	}
}

 */