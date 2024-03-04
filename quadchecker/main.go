package main

import (
	"fmt"
	"io"
	"os"
	"os/exec"
	"strconv"
)

func main() {

	quads := []string{"./quadA", "./quadB", "./quadC", "./quadD", "./quadE"}
	result := []string{}

	data, _ := io.ReadAll(os.Stdin)
	if string(data) == "" {

		fmt.Println("Not a quad function")
		return
	}
	width := 0
	height := 0

	endWidth := true

	for _, one := range data {
		if one == 10 {
			height++
			endWidth = false
		} else if endWidth {
			width++
		}

	}
	for i := 0; i < len(quads); i++ {
		cmd := exec.Command(quads[i], strconv.Itoa(width), strconv.Itoa(height))
		rs, _ := cmd.Output()
		if string(data) == string(rs) {
			result = append(result, string(quads[i]))
		}
	}

	if len(result) == 0 {
		fmt.Println("Not a quad function")
	} else if len(data) >= 1 {
		for i, ch := range result {
			fmt.Printf("[%s] [%d] [%d]", ch[2:], width, height)
			if i < len(result)-1 {
				fmt.Printf(" || ")
			}

		}
		fmt.Println()
	}
}
