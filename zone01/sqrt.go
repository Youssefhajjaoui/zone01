package piscine

func Sqrt(nb int) int {
	rst := 0
	if nb > 0 {
		if nb == 1 {
			return 1
		} else {
			for i := 1; i < nb; i++ {
				if i*i == nb {
					rst = i
				}
			}
		}
	}
	return rst
}
