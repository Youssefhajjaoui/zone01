package piscine

func ConvertBase(nbr, baseFrom, baseTo string) string {
	res := ""
	if check(baseTo) && check(baseFrom) {
		str := AtoiBase(nbr, baseFrom)
		res = PrintNbrBase(str, baseTo)
	}
	return res
}
func check(bs string) bool {
	for i := 0; i < len(bs); i++ {
		for j := i + 1; j < len(bs)-1; j++ {
			if bs[i] == bs[j] {
				return false
			}
		}
	}
	return true
}
