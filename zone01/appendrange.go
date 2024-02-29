package piscine

func AppendRange(min, max int) []int {
	v := max - min
	tab := []int{}

	if v <= 0 {
		return nil
	} else {
		for i := 0; i < v; i++ {
			tab = append(tab, min)
			min++
		}
	}
	return tab
}
