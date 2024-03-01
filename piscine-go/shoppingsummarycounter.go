package piscine

func ShoppingSummaryCounter(str string) map[string]int {
	summary := make(map[string]int)
	word := ""
	for i := 0; i < len(str); i++ {
		if str[i] != ' ' {
			word = word + string(str[i])
		} else if str[i] == ' ' {
			summary[word]++
			word = ""
		}
	}
	if word != " " {
		summary[word]++
	}
	return summary
}
