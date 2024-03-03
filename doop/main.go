package main

import (
	"os"
)

func main() {
	arr := os.Args[1:]
	if len(arr) == 3 {
		a, j := Atoi(os.Args[1])
		b, k := Atoi(os.Args[3])
		if j && k {
			if os.Args[2] == "+" {
				if a < 2147483647 && a > -2147483647 && b < 2147483647 && b > -2147483647 {
					writeInt(a + b)
					PrintString("\n")
				}
			} else if os.Args[2] == "-" {
				if a < 2147483647 && a > -2147483647 && b < 2147483647 && b > -2147483647 {
					writeInt(a - b)
					PrintString("\n")
				}
			} else if os.Args[2] == "/" {
				if a < 2147483647 && a > -2147483647 && b < 2147483647 && b > -2147483647 {
					if b == 0 {
						PrintString("No division by 0\n")
					} else if a >= b {
						writeInt(a / b)
						PrintString("\n")
					} else if a < b {
						PrintString("0\n")
					}
				}
			} else if os.Args[2] == "%" {
				if a < 2147483647 && a > -2147483647 && b < 2147483647 && b > -2147483647 {
					if b == 0 {
						PrintString("No modulo by 0\n")
					} else {
						writeInt(a % b)
						PrintString("\n")
					}
				}
			} else if os.Args[2] == "*" {
				if a < 2147483647 && a > -2147483647 && b < 2147483647 && b > -2147483647 {
					writeInt(a * b)
					PrintString("\n")
				}
			}
		}
	}
}

func PrintString(s string) {
	os.Stdout.WriteString(s)
}

func writeInt(i int) {
	var buf [20]byte
	pos := len(buf)
	isprintable := false
	if i == 0 {
		pos--
		buf[pos] = '0'
	} else {
		neg := false
		if i < 0 && i > -9223372036854775807 {
			neg = true
			i = -i
			isprintable = true
		}
		for i > 0 && i <= 9223372036854775807 {
			pos--
			buf[pos] = byte('0' + i%10)
			i /= 10
			isprintable = true
		}
		if neg {
			pos--
			buf[pos] = '-'
		}
	}
	if isprintable {
		PrintString(string(buf[pos:]))
	}
}

func Atoi(s string) (int, bool) {
	number := 0
	t := true
	if len(s) > 0 {
		factor := 1
		for i := len(s) - 1; i >= 0; i-- {
			if (s[i] < '0' || s[i] > '9') && (i != 0 && (s[0] != '-' && s[0] != '+')) {
				t = false
			}
			if s[i] != '-' && s[i] != '+' {
				number += (int(s[i]) - 48) * factor
				factor = factor * 10
			}
		}
		if s[0] == '-' {
			return -number, t
		}
	}
	return number, t
}
