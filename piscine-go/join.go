package piscine

func Join(strs []string, sep string) string {
	result := ""
	for i, v := range strs {
		if i == len(strs)-1 {
			result += v
		} else if i != len(strs) {
			result += v + sep
		}
	}
	return result
}
