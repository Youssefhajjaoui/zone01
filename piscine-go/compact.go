package piscine

func Compact(ptr *[]string) int {
	t := []string{}
	result := 0
	arr := *ptr
	for i := 0; i < len(arr); i++ {
		if arr[i] != "" {
			t = append(t, arr[i])
			result++
		}
	}
	*ptr = t
	return result
}
