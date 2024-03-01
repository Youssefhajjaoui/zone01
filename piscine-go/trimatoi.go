package piscine

func TrimAtoi(s string) int {
	tab := []rune(s)
	sgn := 1
	result := 0

	for i := 0; i < len(tab); i++ {
		if tab[i] == '-' {
			for j := 0; j < len(tab[0:i]); j++ {
				if tab[j] >= '0' && tab[j] <= '9' {
					sgn = 1
				} else {
					sgn = -1
				}
			}
		} else if IsNumer(string(tab[i])) {
			result = int(tab[i]-'0') + result*10
		}
	}

	return result * sgn
}

func IsNumer(s string) bool {
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
