package piscine

func Index(s string, toFind string) int {
	rs := []rune(s)

	rt := []rune(toFind)
	if len(rt) == 0 {
		return 0
	} else {
		for i := 0; i <= len(rs)-len(rt); i++ {
			rsl := true
			for j := 0; j < len(rt); j++ {
				if rs[i+j] != rt[j] {
					rsl = false
					break
				}
			}
			if rsl == true {
				return i
			}
		}
	}
	return -1
}
