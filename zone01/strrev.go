package piscine

func StrRev(s string) string {
	l := []rune(s)

	str := ""
	for i := len(l) - 1; i >= 0; i-- {
		str += string(l[i])
	}
	return str
}
