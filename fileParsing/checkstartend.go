package lem_in

func CheckStartEnd(str string, indice int, input []string) bool {
	if indice == len(input)-1 {
		return false
	}
	if !CheckRoom(input[indice+1]) {
		return false
	}

	if str == "##start" {
		if Started {
			return false
		} else {
			StartRoom = input[indice+1]
			Started = true
			return true
		}
	} else {
		if Ended {
			return false
		} else {
			EndRoom = input[indice+1]
			Ended = true
			return true
		}
	}
}
