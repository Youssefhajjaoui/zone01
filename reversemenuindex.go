package piscine

func ReverseMenuIndex(menu []string) []string {
	tab:=[]string{}
	for i:=len(menu)-1 ; i>=0 ; i-- {
		tab=append(tab,menu[i])
	}
	return tab
}