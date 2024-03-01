package piscine

func IterativePower(nb int, power int) int {
	rst := 1
	if power < 0 {
		return 0
	} else {
		for i := 0; i < power; i++ {
			rst = rst * nb
		}
	}
	return rst
}
