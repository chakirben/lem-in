package ants

import "fmt"

var (
	AntPositions = make(map[int]string)
	AntPaths     = make(map[int][]string)
	RoomsSate    = make(map[string]int)
	PathwithAnts = make(map[int]int)
)

func FindAntJourney(NOA int, end string, start string) {
	for RoomsSate[end] != NOA {
		for i := 1; i <= NOA; i++ {

			Move(i, end, start)
		}
		fmt.Println()
	}
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
	if len(AntPaths[i]) < 2 {
		return false
	}
	// fmt.Println("------------")
	// fmt.Println("the ant is :", i, "currently at  ", AntPaths[i][0], "it's next room is ", AntPaths[i][1], "and it's ", RoomsSate[AntPaths[i][1]])
	// check if next room is end
	//fmt.Println(AntPaths)
	if AntPaths[i][1] == end {
		fmt.Printf("L%v-%v ", i, AntPaths[i][1])
		RoomsSate[AntPaths[i][0]]--
		AntPositions[i] = end
		AntPaths[i] = []string{}
		RoomsSate[end]++
		return true
	}
	if RoomsSate[AntPaths[i][1]] == 0 {
		fmt.Printf("L%v-%v ", i, AntPaths[i][1])
		RoomsSate[AntPositions[i]]--
		AntPositions[i] = AntPaths[i][1]
		RoomsSate[AntPaths[i][1]]++
		AntPaths[i] = AntPaths[i][1:]
		return true
	}

	/* if AntPaths[i][0] == start {
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
	// move to room if it's empty */
	return false
}

func SpreadAnts(NOA int, paths [][]string) {

	indice := 0
	Sindice := 1
	for i := 1; i <= NOA; i++ {
		if len(paths) == 1 {
			fmt.Println(i)
			AntPaths[i+1] = paths[0]
			continue
		}
		fmt.Println("ant ", i, " indice is ", indice, " to compaire is ", Sindice)
		// gives the first ant the first path
		fmt.Println(PathwithAnts)
		indice = min(PathwithAnts)
		AntPaths[i] = paths[indice]
		PathwithAnts[indice]++
		fmt.Println(indice)
		// if i == 1 {
		// 	AntPaths[i] = paths[0]
		// 	PathwithAnts[indice]++
		// 	continue
		// }
		// // if  you're in the middle  compare which path is shorter
		// if len(paths)-1 > indice && indice > 0 {
		// 	fmt.Println("compared with ", Sindice)
		// 	Sindice = Getshorteindice(indice-1, indice+1)
		// }
		// // if you are at the end compaire with the previous
		// if len(paths)-1 == indice {
		// 	if indice > 0 {
		// 		Sindice = indice - 1
		// 	}
		// }
		// // compaire with the shotest path arround to swap
		// if PathwithAnts[indice] > PathwithAnts[Sindice] {
		// 	indice = Sindice
		// }
		// // if the current path didn't change it will keep it
		// AntPaths[i] = paths[indice]
		// PathwithAnts[indice]++

	}
}

func min(m map[int]int) int {
	fmt.Println("Test Path in Min: ", m)
	theMin := m[0]
	mi := 0
	for key, value := range m {
		if theMin > value {
			theMin = value
			mi = key
		}
	}
	return mi
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
