package piscine

func Any(f func(string) bool, a []string) bool {
	var t bool

	for _, v := range a {
		t = f(v)
		if t == true {
			break
		}
	}
	return t
}
