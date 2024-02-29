package piscine

func IsPrintable(s string) bool {
	var up bool
	rs := []rune(s)
	for i := 0; i <= len(s)-1; i++ {
		if rs[i] >= 32 && rs[i] <= 126 {
			up = true
		} else {
			up = false
			break

		}
	}
	return up
}
