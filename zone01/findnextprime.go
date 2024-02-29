package piscine

func FindNextPrime(nb int) int {
	if IsPri(nb) {
		return nb
	} else {
		return FindNextPrime(nb + 1)
	}
}

func IsPri(nb int) bool {
	rs := true
	if nb <= 1 {
		return false
	} else if nb == 2 {
		return true
	} else if nb == 3 {
		return true
	} else if nb%2 == 0 {
		return false
	} else if nb%3 == 0 {
		return false
	} else {
		for i := 3; i*i <= nb; i += 2 {
			if nb%i == 0 {
				rs = false
				break
			}
		}
	}
	return rs
}
