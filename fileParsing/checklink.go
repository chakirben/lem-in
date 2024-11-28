package lem_in

import "strings"

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
	if !RoomExist(arr[0], Rooms) || !RoomExist(arr[1], Rooms) {
		return false
	}
	if IsRepeted(str) {
		return false
	}
	return true
}

func RoomExist(str string, Rooms []string) bool {
	for _, ele := range Rooms {
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
