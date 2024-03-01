package piscine

func DescendAppendRange(max, min int) []int {
	tab := []int{}
	if max < min {
		return tab
	} else {
		for i := max; i > min; i-- {
			tab = append(tab, i)
		}
	}
	return tab
}
