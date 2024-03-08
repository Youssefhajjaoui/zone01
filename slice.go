package piscine

func Slice(a []string, nbrs ...int) []string {
	if len(nbrs) == 0 {
		return a
	}
	if len(nbrs) == 1 {
		first := nbrs[0]
		if first < 0 {
			first = len(a) + first
			if first < 0 {
				return a
			}
		}
		return a[first:]

	}
	first := nbrs[0]
	second := nbrs[1]
	first = covetops(first, a)
	second = covetops(second, a)
	if first > second {
		return nil
	}
	return a[first:second]
}
func covetops(nbr int, a []string) int {
	if nbr < 0 {
		nbr = len(a) + nbr
	}
	if nbr < 0 {
		nbr = 0
	} else if nbr > len(a) {
		nbr = len(a)
	}
	return nbr
}
