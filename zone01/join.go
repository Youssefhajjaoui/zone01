package piscine

func Join(strs []string, sep string) string {
	rs := []string(strs)
	var st string
	for i := 0; i < len(rs); i++ {
		for j := 0; j < len(strs); j++ {
			st += strs[j]
			if j != len(rs)-1 {
				st += sep
			}
		}
		break
	}
	return st
}
