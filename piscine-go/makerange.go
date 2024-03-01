package piscine

func MakeRange(min, max int) []int {
	tab := []int{}
	if max-min > 0 {
		tab = make([]int, max-min)
		t := max - min

		for i := 0; i < t; i++ {
			tab[i] = min
			min++
		}
	} else {
		return nil
	}
	return tab
}
