package piscine

func CountIf(f func(string) bool, tab []string) int {
	var t bool
	var v int
	for _, k := range tab {
		t = f(k)
		if t == true {
			v++
		}
	}
	return v
}
