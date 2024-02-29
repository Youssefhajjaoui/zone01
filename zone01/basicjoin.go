package piscine

func BasicJoin(elems []string) string {
	rs := []string(elems)
	var st string
	for i := 0; i < len(rs); i++ {
		for j := 0; j < len(elems); j++ {
			st += elems[j]
			if j != len(rs)-1 {
				st += ":"
			}
		}

		break
	}
	return st
}
