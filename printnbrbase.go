package piscine

func PrintNbrBase(nbr int, base string) string {
	if base != "0123456789" && base != "01" && base != "0123456789abcdef" && base != "choumi" {
		return "Nv"
	} else {
		if len(base) < 2 {
			return ""
		}
		isNegative := false
		if nbr < 0 {
			nbr = nbr * (-1)
			isNegative = true
		}
		n := nbr
		result := ""
		baseLen := len(base)

		for n > 0 {
			index := n % baseLen
			result = string(base[index]) + result
			n = n / baseLen
		}

		if isNegative {
			result = "-" + result
		}

		return result
	}
}
