package ants

import "strings"

func FindAntJourney(paths []string) {
	j := [][]string{}
	for _, h := range paths {
		j = append(j, strings.Split(h, "-"))
	}


}
func IsValidStep() {

}
/*how to do this 
in my mind i well do somthing lik range for all path and check is valid step for the ants
if this is good idea  i dont now but i work on it if i have a problem in my logic tell me 
*/



/*explain the logic */

/*
the function check step by step if the step is valid the ant go in this roms but if not valid the ant break one step
*/