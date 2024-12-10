package algo

import (
	"fmt"
	"strconv"

	lem_in "lem_in/fileParsing"
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
			fmt.Println(ways)
			break
		}
	}
}
