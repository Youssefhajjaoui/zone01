package piscine

func StringToIntSlice(str string) []int {
	tab := []int{}
	if len(str) == 0 {
		tab = nil
	} else {
		for _, v := range str {
			tab = append(tab, int(v))
		}
	}
	return tab
}
