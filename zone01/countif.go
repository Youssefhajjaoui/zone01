package piscine

func CountI(f func(string) bool, tab []string) int {
	var t int
	for _,v :=range tab {
		if f(v)==true{
			t++
		}
	}
	return t
}
