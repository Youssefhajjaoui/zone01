package piscine

func Map(f func(int) bool, a []int) []bool {
	tab := []bool{}
	for _, v := range a {
		tab = append(tab, f(v))
	}
	return tab
}
