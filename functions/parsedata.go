package lem_in

import (
	"fmt"
	"strconv"
)

func ParseFarm(input []string) (NumberOfAnts int, rooms []string, links []string, err error) {
	DeleteComments(input)
	NOA, err2 := CheckNOfAnts(input[0])
	if err2 != nil {
		return NOA, nil, nil, err2
	}
	for i, line := range input {
		
	}
}

func CheckNOfAnts(line1 string) (int, error) {
	NumberOfAnts, err := strconv.Atoi(line1)
	if err != nil {
		return 0, fmt.Errorf("no Number of ants")
	}
	return NumberOfAnts, nil
}
