package main

import (
	"os"

	"github.com/01-edu/z01"
)

func PrintStr(str string) {
	for _, c := range str {
		z01.PrintRune(c)
	}
	z01.PrintRune('\n')
}

func main() {
	if len(os.Args) == 2 {
		str := os.Args[1]
		if str[0] == 'a' || str[0] == 'e' || str[0] == 'i' || str[0] == 'o' || str[0] == 'u' || str[0] == 'A' || str[0] == 'E' || str[0] == 'I' || str[0] == 'O' || str[0] == 'U' {
			PrintStr(str + "ay")
		} else {
			find := -1 // Initialize find to -1 to indicate no consonants found yet
			for i := 0; i < len(str); i++ {
				if str[i] == 'a' || str[i] == 'e' || str[i] == 'i' || str[i] == 'o' || str[i] == 'u' || str[i] == 'A' || str[i] == 'E' || str[i] == 'I' || str[i] == 'O' || str[i] == 'U' {
					if find == -1 {
						find = i
					}
				}
			}
			// If find is still -1, it means the word has no consonants (only a vowel)
			if find == -1 {
				print("No vowels")
			} else {
				PrintStr(str[find:] + str[:find] + "ay")
			}
		}
	}
}
