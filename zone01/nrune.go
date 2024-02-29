package piscine

func NRune(s string, n int) rune {
	var fnt rune
	rs := []rune(s)
	if n > 0 && n <= len(rs) {
		for i := 0; i <= len(rs); i++ {
			if i == n-1 {
				fnt = rs[i]
			}
		}
	}
	return fnt
}
