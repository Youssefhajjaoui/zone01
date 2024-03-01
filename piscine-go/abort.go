package piscine

func Abort(a, b, c, d, e int) int {
	tr := []int{a, b, c, d, e}
	for i := 0; i <= 4; i++ {
		for j := i + 1; j <= 4; j++ {
			if tr[i] < tr[j] {
				tr[i], tr[j] = tr[j], tr[i]
			}
		}
	}
	return tr[2]
}
