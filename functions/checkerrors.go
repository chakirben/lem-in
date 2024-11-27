package lem_in

import (
	"strconv"
	"strings"
)

func CheckEroors(arr []string) bool {
	if len(arr) == 0 {
		return false
	}
	arr = DeleteComments(arr)
	CheckNOfAnts(arr)
	CheckStarEnd(arr)
	return true
}

func CheckNOfAnts(arr []string) bool {
	_, err := strconv.Atoi(arr[0])
	return err != nil
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
