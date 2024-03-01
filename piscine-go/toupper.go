package piscine

func ToUpper(s string) string {
	result := ""
	rs := []rune(s)
	for i := 0; i <= len(s)-1; i++ {
		if rs[i] >= 'a' && rs[i] <= 'z' {
			result = result + string(rs[i]-32)
		} else {
			result = result + string(rs[i])
		}
	}
	return result
}
