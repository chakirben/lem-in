package ants

import "fmt"

var (
	AntPositions = make(map[int]string)
	AntPaths     = make(map[int][]string)
	RoomsSate    = make(map[string]int)
	PathwithAnts = make(map[int]int)
)

func FindAntJourney(NOA int, end string, start string) {
	if RoomsSate[end] == NOA {
		return
	}
	for i := range NOA {
		if Move(i, end, start) {
			if len(AntPaths[i]) == 0 {
				fmt.Printf("L%v-%v ", AntPositions[i], end)
				continue
			}
			fmt.Printf("L%v-%v ", AntPositions[i], AntPaths[i][0])
		}
	}
	fmt.Println()
	FindAntJourney(NOA, end, start)
}

func DeleteStart() {
	for a, p := range AntPaths {
		AntPaths[a] = p[1:]
	}
}

func InitRooms(rooms []string, start string, NOA int) {
	for _, r := range rooms {
		RoomsSate[r] = 0
	}
	RoomsSate[start] = NOA
}

func InitPathlength(paths [][]string) {
	for i, pa := range paths {
		PathwithAnts[i] = len(pa)
	}
}

func Move(i int, end string, start string) bool {
	// check if the ant has already arrived !
	if len(AntPaths[i]) == 0 {
		return false
	}
	if AntPaths[i][0] == start {
		AntPositions[i] = start
		AntPaths[i] = AntPaths[i][1:]
		RoomsSate[start]++
		return true
	}
	// check if next room is end
	if AntPaths[i][0] == end {
		AntPositions[i] = end
		AntPaths[i] = []string{}
		RoomsSate[end]++
		return true
	}
	// move to room if it's empty
	if RoomsSate[AntPaths[i][0]] == 0 {
		RoomsSate[AntPositions[i]]--
		AntPositions[i] = AntPaths[i][0]
		RoomsSate[AntPaths[i][0]]++
		AntPaths[i] = AntPaths[i][1:]
		return true
	}
	return false
}

func SpreadAnts(NOA int, paths [][]string) {
	indice := 0
	Sindice := 1
	for i := range NOA {
		// gives the first ant the first path
		fmt.Println(PathwithAnts)
		if i == 0 {
			AntPaths[i] = paths[0]
			PathwithAnts[indice]++
			continue
		}
		// if  you're in the middle  compare which path is shorter
		if len(paths)-1 > indice && indice > 0 {
			Sindice = Getshorteindice(indice-1, indice+1)
		}
		// if you are at the end compaire with the previous
		if len(paths)-1 == indice {
			if indice > 0 {
				Sindice = indice - 1
			}
		}
		// compaire with the shotest path arround to swap
		if PathwithAnts[indice] > PathwithAnts[Sindice] {
			indice = Sindice
		}
		// if the current path didn't change it will keep it
		AntPaths[i] = paths[indice]
		PathwithAnts[indice]++

	}
}

func Getshorteindice(i int, j int) int {
	if PathwithAnts[i] > PathwithAnts[j] {
		return j
	}
	return i
}

func InitPositions(s string, NOA int) {
	for a := range NOA {
		AntPositions[a] = s
	}
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
/* PathIndice := 0
CurrentPath := paths[PathIndice]
Pathlength := len(CurrentPath)
prev := 0
next := 0
current := len(paths[0])
for i := range NOA {
	if i == 0 {
		AntPaths[i+1] = paths[PathIndice]
		current++
	}
	if len(paths)-1 > PathIndice && Pathlength > len(paths[PathIndice+1]) {
		prev = current
		current = len()
		CurrentPath = paths[PathIndice+1]
		fmt.Println("ay")
		PathIndice++
	}
	AntPaths[i+1] = CurrentPath
	Pathlength++
} */
