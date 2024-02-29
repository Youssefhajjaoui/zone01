package piscine

func IsAlpha(s string) bool {
	var up bool
	rs := []rune(s)
	for i := 0; i <= len(s)-1; i++ {
		if rs[i] >= 'a' && rs[i] <= 'z' || rs[i] >= 'A' && rs[i] <= 'Z' || rs[i] >= '0' && rs[i] <= '9' {
			up = true
		} else {
			up = false
			break

		}
	}
	return up
}
