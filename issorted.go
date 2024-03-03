package piscine

func IsSorted(f func(a, b int) int, a []int) bool {
	if len(a) > 1 {
		if f(a[0], a[1]) > 0 {
			for i, j := 0, 1; i < len(a)-1 && j < len(a); i++ {
				if f(a[i], a[j]) < 0 {
					return false
				}
				j++
			}
		} else if f(a[0], a[1]) < 0 {
			for i, j := 0, 1; i < len(a)-1 && j < len(a); i++ {
				if f(a[i], a[j]) > 0 {
					return false
				}
				j++
			}
		}
	}
	return true
}
