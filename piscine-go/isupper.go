package piscine

func IsUpper(s string) bool {
	var up bool
	rs := []rune(s)
	for i := 0; i <= len(s)-1; i++ {
		if rs[i] >= 'A' && rs[i] <= 'Z' {
			up = true
		} else {
			up = false
			break

		}
	}
	return up
}
