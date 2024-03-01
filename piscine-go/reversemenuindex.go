package piscine

func ReverseMenuIndex(menu []string) []string {
	for i := 0; i < len(menu)/2; i++ {
		menu[i], menu[len(menu)-i-1] = menu[len(menu)-i-1], menu[i]
	}
	return menu
}
