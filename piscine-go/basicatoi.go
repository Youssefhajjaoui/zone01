package piscine

func BasicAtoi(s string) int {
	res := 0
	str := []rune(s)
	var cnv int
	for i := 0; i < len(str); i++ {
		if str[i] < 48 || str[i] > 57 {
			return -1
		}
		cnv = int(str[i] - 48)
		res = res*10 + cnv
	}

	return res
}
