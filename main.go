package main

import (
	"fmt"
	algo "lem-in/farm"
	"lem-in/server"
	"os"
	"time"
)

func main() {

	CurrentTime := time.Now()
	fmt.Println("Time Start: ", CurrentTime.Format("15:04:05"))
	if len(os.Args) != 2 {
		PrintArgError()
		return
	}
	filepath := os.Args[1]
	err := server.ReadFile(filepath)
	if err != "" {
		fmt.Println(err)
		return
	}
	//server.RoomChecks()
	server.PrintAll()

	algo.BuildFarmFromRooms()
	fmt.Println("FarmBuilt at: ", CurrentTime.Format("15:04:05"))

	// farm := algo.Farm
	// for k, v := range farm {
	// 	fmt.Println("Room: ", k, "Linked to the following: ", v)
	// }
	algo.MaxFlow() // max numver of paths
	fmt.Println("MaxFlow Generated at:", CurrentTime.Format("15:04:05"))

	//fmt.Println("Max number of paths: ", maxNumPaths)
	paths := algo.ExtractPaths(algo.Farm)

	algo.AntMovement(paths)
	fmt.Println("all ants reached End safely at: ", CurrentTime.Format("15:04:05"))
	fmt.Println("max Flows", algo.MF)
	fmt.Println("max Flows", algo.MF, len(paths))

}

func PrintArgError() {
	fmt.Println("Invalid number of arguments")
	fmt.Println("Usage: go run main.go <file_path>")
}
