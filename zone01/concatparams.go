package piscine

func ConcatParams(args []string) string {
	// tab := []string{}
	st := ""
	for i := 0; i < len(args); i++ {
		// tab := append(tab, args[i])
		if i == len(args)-1 {
			st += args[i]
		} else {
			st += args[i] + "\n"
		}
	}
	return st
}
