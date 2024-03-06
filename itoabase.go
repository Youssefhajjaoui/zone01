package piscine

func ItoaBase(value, base int) string {
	baseorigin := "0123456789ABCDEF"
	newbase := make(map[int]string)
	st := ""
	for i, v := range baseorigin {
		newbase[i] = string(v)
	}
	if value < 0 {
		for value != 0 {
			digit := value % base
			if digit < 0 {
				digit = -digit
			}
			st = newbase[digit] + st
			value = value / base
		}
		return "-" + st
	}
	if value > 0 {
		for value != 0 {
			digit := value % base
			st = newbase[digit] + st
			value = value / base
		}
	}
	return st
}
