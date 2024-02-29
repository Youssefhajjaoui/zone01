package piscine

func IsLower(s string) bool {
	var up bool
	rs := []rune(s)
	for i := 0; i <= len(s)-1; i++ {
		if rs[i] >= 'a' && rs[i] <= 'z' {
			up = true
		} else {
			up = false
			break

		}
	}
	return up
}
