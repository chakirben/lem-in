package lem_in

import (
	"fmt"
	"strconv"
	"strings"
)

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
			s := strings.Split(input[indice+1], " ")
			Fa.Start = s[0]
			Started = true
			return true
		}
	} else {
		if Ended {
			return false
		} else {
			e := strings.Split(input[indice+1], " ")
			Fa.End = e[0]
			Ended = true
			return true
		}
	}
}

func CheckRoom(str string) bool {
	arr := strings.Split(str, " ")
	if len(arr) != 3 {
		return false
	}
	_, err := strconv.Atoi(arr[1])
	if err != nil {
		return false
	}
	_, err2 := strconv.Atoi(arr[2])
	if err2 != nil {
		return false
	}
	return true
}

func IsCooRepeated(str string) bool {
	count := 0
	for _, coo := range Coordinations {
		if coo == str {
			count++
		}
	}
	return count > 1
}

func IsRoomRepeated(room string) bool {
	for room2 := range Fa.Rooms {
		if room == room2 {
			return true
		}
	}
	return false
}

func Checklink(str string) bool {
	if len(str) < 3 {
		return false
	}
	// hadi bach kanchof wach kayna - f wst 2 strings
	arr := strings.Split(str, "-")
	if len(arr) != 2 {
		return false
	}
	// hadi bach nchof dok rooms wach kaynin
	if !RoomExist(arr[0], Fa.Rooms) || !RoomExist(arr[1], Fa.Rooms) {
		return false
	}
	if IsRepeted(str) {
		return false
	}
	return true
}

func RoomExist(str string, Rooms map[string][]string) bool {
	for ele := range Rooms {
		if ele == str {
			return true
		}
	}
	return false
}

func IsRepeted(str string) bool {
	swapedLink := Swap(str)
	count := 0
	for _, link := range Links {
		if link == str {
			count++
		}
		if link == swapedLink {
			return true
		}
	}
	return count > 1
}

func Swap(str string) string {
	str1 := ""
	arr := strings.Split(str, "-")
	str1 = arr[1] + "-" + arr[0]
	return str1
}

func CheckNOfAnts(line1 string) error {
	NumberOfAnts, err := strconv.Atoi(line1)
	Fa.NOA = NumberOfAnts
	if err != nil {
		return fmt.Errorf("no Number of ants")
	}
	return nil
}

func DeleteComments(arr []string) []string {
	arr2 := []string{}
	for _, str := range arr {
		if str == "##start" || str == "##end" {
			arr2 = append(arr2, str)
			continue
		}
		if str == "" {
			continue
		}
		if str[0] == '#' {
			continue
		}
		arr2 = append(arr2, str)
	}
	return arr2
}
