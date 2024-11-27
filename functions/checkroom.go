package lem_in

import (
	"strconv"
	"strings"
)

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
	return err2 == nil
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
	for _, room2 := range Rooms {
		if room == room2 {
			return true
		}
	}
	return false
}
