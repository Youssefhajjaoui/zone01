package piscine

func CollatzCountdown(start int) int {
	if start <= 0 {
		return -1
	}
	j := 0
	for start > 1 {
		if start%2 == 0 {
			start = start / 2
			j++
		} else if start%2 != 0 {
			start = start*3 + 1
			j++
		}
	}
	return j
}
