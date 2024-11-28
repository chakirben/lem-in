package lem_in

import (
	"fmt"
	"strconv"
)

func CheckNOfAnts(line1 string) (int, error) {
	NumberOfAnts, err := strconv.Atoi(line1)
	if err != nil {
		return 0, fmt.Errorf("no Number of ants")
	}
	return NumberOfAnts, nil
}
