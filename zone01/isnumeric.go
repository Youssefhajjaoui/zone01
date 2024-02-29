package piscine

func IsNumeric(s string) bool {
	var up bool
	rs := []rune(s)
	for i := 0; i <= len(s)-1; i++ {
		if rs[i] >= '0' && rs[i] <= '9' {
			up = true
		} else {
			up = false
			break

		}
	}
	return up
}
