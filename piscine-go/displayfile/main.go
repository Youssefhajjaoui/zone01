package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	filename := os.Args[1:]
	if len(filename) < 1 {
		fmt.Println("File name missing")
	} else if len(filename) > 1 {
		fmt.Println("Too many arguments")
	} else {
		result := ""
		for _, v := range filename {
			result += v
		}
		content, err := ioutil.ReadFile(result)
		if err != nil {
			fmt.Println("error")
		} else {
			fmt.Printf("%s", content)
		}
	}
}
