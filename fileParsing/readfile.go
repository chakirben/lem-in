package lem_in

import (
	"bufio"
	"log"
	"os"
)

func Readfile() []string {
	file, err := os.Open("inputfiles/" + os.Args[1])
	if err != nil {
		log.Fatal(err)
	}
	scanner := bufio.NewScanner(file)
	arr := []string{}
	for scanner.Scan() {
		arr = append(arr, scanner.Text())
	}
	return arr
}
