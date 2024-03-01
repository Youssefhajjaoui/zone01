package piscine

func JumpOver(str string) string {
	result := ""
	for i := 0; i < len(str); i++ {
		if (i+1)%3 == 0 {
			result += string(str[i])
		}
	}
	result += "\n"
	return result
}
