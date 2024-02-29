package piscine

func AlphaCount(s string) int {
	rs := []rune(s)
	j := 0
	for i := 0; i < len(rs); i++ {
		if rs[i] >= 'a' && rs[i] <= 'z' || rs[i] >= 'A' && rs[i] <= 'Z' {
			j++
		}
	}
	return j
}
