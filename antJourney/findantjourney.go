package algo

import (
	"fmt"
	lem_in "lem_in/fileParsing"
	"strconv"
)

func ini(s [][]string) map[int]int {
	mapp := make(map[int]int)
	for v := range s {
		mapp[v] = len(s[v])
	}
	return mapp
}

func getmin(v map[int]int) int {
	i := v[0]
	x := 0
	for t, vv := range v {
		if vv < i {
			i = vv
			x = t
		}
	}
	return x
}

func initt(i int) [][]string {
	ff := [][]string{}
	for t := 0; t < i; t++ {
		ff = append(ff, []string{})
	}
	return ff
}

func Sendants(ways [][]string) {
	antgroups := initt(len(ways))
	antid := 1
	sss := ini(ways)

	for antid <= lem_in.Fa.Ants {
		s := getmin(sss)
		antgroups[s] = append(antgroups[s], "L"+strconv.Itoa(antid))
		antid++
		sss[s]++
	}
	controltrafic(antgroups, ways)
}

func controltrafic(antgroups, ways [][]string) {
	trafic := make(map[string]int)
	unavailablerooms := make(map[string]bool)
	finished := []string{}
	stallCount := 0
	maxStalls := 100

	for len(finished) != lem_in.Fa.Ants {
		progress := false

		for i := 0; i < len(ways); i++ {
			unavailablerooms[lem_in.Fa.End] = false

			for s := 0; s < len(antgroups[i]); s++ {
				ant := antgroups[i][s]
				if trafic[ant]+1 < len(ways[i]) && !unavailablerooms[ways[i][trafic[ant]+1]] {
					if ways[i][trafic[ant]+1] == lem_in.Fa.End {
						unavailablerooms[ways[i][trafic[ant]]] = false
						finished = append(finished, ant)
						delete(trafic, ant)
						antgroups[i] = append(antgroups[i][:s], antgroups[i][s+1:]...)
						fmt.Printf("%v-%v ", ant, lem_in.Fa.End)
						s--
						unavailablerooms[lem_in.Fa.End] = true
						progress = true
					} else {
						fmt.Printf("%v-%v ", ant, ways[i][trafic[ant]+1])
						unavailablerooms[ways[i][trafic[ant]+1]] = true
						unavailablerooms[ways[i][trafic[ant]]] = false
						trafic[ant]++
						progress = true
					}
				}
			}
		}

		fmt.Println()

		if progress {
			stallCount = 0
		} else {
			stallCount++
		}

		if stallCount > maxStalls {
			break
		}
	}
}

// another way

// package algo

// import (
// 	"fmt"
// 	lem_in "lem_in/fileParsing"
// 	"strconv"
// 	"strings"
// )

// // Helper functions remain the same
// func ini(s [][]string) map[int]int {
// 	mapp := make(map[int]int)
// 	for v := range s {
// 		mapp[v] = len(s[v])
// 	}
// 	return mapp
// }

// func getmin(v map[int]int) int {
// 	i := v[0]
// 	x := 0
// 	for t, vv := range v {
// 		if vv < i {
// 			i = vv
// 			x = t
// 		}
// 	}
// 	return x
// }

// func initt(i int) [][]string {
// 	ff := [][]string{}
// 	for t := 0; t < i; t++ {
// 		ff = append(ff, []string{})
// 	}
// 	return ff
// }

// // Collects the result of Sendants into a [][]string
// func Sendants(ways [][]string) [][]string {
// 	antgroups := initt(len(ways))
// 	antid := 1
// 	sss := ini(ways)
// 	var result []string

// 	for antid <= lem_in.Fa.Ants {
// 		s := getmin(sss)
// 		antgroups[s] = append(antgroups[s], "L"+strconv.Itoa(antid))
// 		antid++
// 		sss[s]++
// 	}
// 	controltrafic(antgroups, ways, &result)

// 	// Split result into [][]string format
// 	var finalResult [][]string
// 	for _, line := range result {
// 		finalResult = append(finalResult, strings.Fields(line))
// 	}
// 	return finalResult
// }

// // Modified to append the printed results into the `output` slice
// func controltrafic(antgroups, ways [][]string, output *[]string) {
// 	trafic := make(map[string]int)
// 	unavailablerooms := make(map[string]bool)
// 	finished := []string{}
// 	stallCount := 0
// 	maxStalls := 100

// 	for len(finished) != lem_in.Fa.Ants {
// 		progress := false
// 		var currentLine []string

// 		for i := 0; i < len(ways); i++ {
// 			unavailablerooms[lem_in.Fa.End] = false

// 			for s := 0; s < len(antgroups[i]); s++ {
// 				ant := antgroups[i][s]
// 				if trafic[ant]+1 < len(ways[i]) && !unavailablerooms[ways[i][trafic[ant]+1]] {
// 					if ways[i][trafic[ant]+1] == lem_in.Fa.End {
// 						unavailablerooms[ways[i][trafic[ant]]] = false
// 						finished = append(finished, ant)
// 						delete(trafic, ant)
// 						antgroups[i] = append(antgroups[i][:s], antgroups[i][s+1:]...)
// 						currentLine = append(currentLine, fmt.Sprintf("%v-%v", ant, lem_in.Fa.End))
// 						s--
// 						unavailablerooms[lem_in.Fa.End] = true
// 						progress = true
// 					} else {
// 						currentLine = append(currentLine, fmt.Sprintf("%v-%v", ant, ways[i][trafic[ant]+1]))
// 						unavailablerooms[ways[i][trafic[ant]+1]] = true
// 						unavailablerooms[ways[i][trafic[ant]]] = false
// 						trafic[ant]++
// 						progress = true
// 					}
// 				}
// 			}
// 		}

// 		// Append the line to the output
// 		if len(currentLine) > 0 {
// 			*output = append(*output, strings.Join(currentLine, " "))
// 		}

// 		if progress {
// 			stallCount = 0
// 		} else {
// 			stallCount++
// 		}

// 		if stallCount > maxStalls {
// 			break
// 		}
// 	}
// }

// // New function to compare runs
// func CompareRuns(ways [][]string, runs int) (map[int][][]string, [][]string) {
// 	results := make(map[int][][]string)
// 	minSteps := int(^uint(0) >> 1) // Initialize to max int
// 	var bestRun [][]string

// 	for i := 0; i < runs; i++ {
// 		result := Sendants(ways)
// 		results[i] = result

// 		if len(result) < minSteps {
// 			minSteps = len(result)
// 			bestRun = result
// 		}
// 	}

// 	return results, bestRun
// }
