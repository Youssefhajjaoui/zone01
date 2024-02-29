package piscine

func SplitWhiteSpaces(s string) []string {
	res := []string{}
	tab := []rune(s)
	find := 0
	// tes := []string{}
	for i := 0; i < len(tab); i++ {
		if tab[i] == ' ' || tab[i] == '\n' {
			if i != find {
				res = append(res, string(tab[find:i]))
			}
			find = i + 1
		} else if i == len(s) {

			res = append(res, string(tab[find:i]))
			find = i + 1
		}
	}

	res = append(res, string(tab[find:]))

	// if len(res[0]) == 0 {
	// 	res = append(res[:0], res[1:]...)
	// }
	return res
}
