package lem_in

import (
	"fmt"
	"strconv"
)
var StartRoom string
var Started bool = false
func ParseFarm(input []string) (NumberOfAnts int, rooms []string, links []string, err error) {
	DeleteComments(input)
	NOA, err2 := CheckNOfAnts(input[0])
	if err2 != nil {
		return NOA, nil, nil, err2
	}
	for i, line := range input {
		if i == 0 {
			continue
		}
		if line == "start" {
			if CheckStart(line, i, input) {
				StartRoom = line
				Started = true
			} else {
				return NOA, nil, nil, err2
			}
		}
		if line == "end" {
			if CheckEnd(line, i, input) {
				StartRoom = line
				Started = true
			} else {
				return NOA, nil, nil, err2
			}
		}
		

	}
}

func CheckStart(str string, indice int, input []string) bool {
	if indice == len(input)-1 {
		return false
	}
	if Started {
		return false
	}
	if CheckRoom(input[indice+1]) {
		return true
	}
	return false
}
func CheckEnd(str string, indice int, input []string) bool {
	if indice == len(input)-1 {
		return false
	}
	if Started {
		return false
	}
	if CheckRoom(input[indice+1]) {
		return true
	}
	return false
}

func CheckNOfAnts(line1 string) (int, error) {
	NumberOfAnts, err := strconv.Atoi(line1)
	if err != nil {
		return 0, fmt.Errorf("no Number of ants")
	}
	return NumberOfAnts, nil
}
