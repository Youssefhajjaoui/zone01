package piscine

func ToLower(s string) string {
	result := ""
	rs := []rune(s)
	for i := 0; i <= len(s)-1; i++ {
		if rs[i] >= 'A' && rs[i] <= 'Z' {
			result = result + string(rs[i]+32)
		} else {
			result = result + string(rs[i])
		}
	}
	return result
}
