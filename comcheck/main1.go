package main

import (
    "fmt"
    "os"
)

func main() {
    tab := os.Args[1:]
    for _, v := range tab {
        if v == "01" || v == "galaxy" || v == "galaxy 01" {
            fmt.Println("Alert!!!")
            break
        }
    }
}