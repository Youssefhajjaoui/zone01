		package piscine

func AtoiBase(s string, base string) int {
	rs := 0
	if !checkifbase(base) {
		return 0 
	}
	baseMap := make(map[byte]int)
	for i := 0; i < len(base); i++ {
		baseMap[base[i]] = i
	}
	for i := len(s) - 1; i >= 0; i-- {
		digit := baseMap[s[i]]
		rs += digit * powercount(len(base), len(s)-1-i)
	}
	return rs
}
func checkifbase(base string) bool {
	if len(base) < 2 {
		return false
	}
	chart := make(map[string]bool)
	for _, v := range base {
		if string(v) == "-" || string(v) == "+" {
			return false
		} else if !chart[string(v)] {
			chart[string(v)] = true
		} else {
			return false
		}
	}
	return true
}
func powercount(v int, power int) int {
	rst := 1
	if power < 0 {
		return 0
	} else {
		for i := 0; i < power; i++ {
			rst = rst * v
		}
	}
	return rst
}
