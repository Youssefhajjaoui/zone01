package piscine

func RecursivePower(nb int, power int) int {
	rsl := 1
	if power < 0 {
		return 0
	} else if power > 0 {
		rsl = nb * RecursivePower(nb, power-1)
	}
	return rsl
}
