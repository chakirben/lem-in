package lem_in

import (
	"strconv"
	"strings"
)

var Rooms []string

func CheckEroors(arr []string) bool {
	if len(arr) == 0 {
		return false
	}
	arr = DeleteComments(arr)
	CheckNOfAnts(arr)
	CheckStarEnd(arr)
	Checklink(arr[1])
	return true
}

func CheckStarEnd(arr []string) bool {
	StartCount := 0
	EndCount := 0
	for i, str := range arr {
		if str == "##start" || str == "##end" {
			if i == len(arr)-1 {
				return false
			}
			if CheckRoom(arr[i+1]) {
				if str == "##start" {
					StartCount++
				} else {
					EndCount++
				}
			} else {
				return false
			}
		}
	}
	return StartCount == 1 && EndCount == 1
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
	_, err2 := strconv.Atoi(arr[1])
	return err2 == nil
}

func DeleteComments(arr []string) []string {
	arr2 := []string{}
	for _, str := range arr {
		if str == "##start" || str == "##finish" {
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
	if !RoomExist(arr[0], Rooms) || !RoomExist(arr[2], Rooms) {
		return false
	}
	return true 
}

func containlink(str string) bool {
	for i, char := range str {
		if i != 0 && char == '-' {
			return true
		}
	}
	return false
}

func RoomExist(str string, Rooms []string) bool {
	for _, ele := range Rooms {
		if ele == str {
			return true
		}
	}
	return false
}
