package lem_in

func DeleteComments(arr []string) []string {
	arr2 := []string{}
	for _, str := range arr {
		if str == "##start" || str == "##end" {
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
