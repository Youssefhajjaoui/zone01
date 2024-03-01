package main

import "github.com/01-edu/z01"
type Door struct{
	state string
}


func PrintStr(s string) {
	for _, r := range s {
		z01.PrintRune(r)
	}
}

func CloseDoor(ptrDoor *Door) {
	PrintStr("Door Closing...")
	ptrDoor.state = "CLOSE"
}
func OpenDoor(ptrDoor *Door) {
	PrintStr("Door OPENING...")
	ptrDoor.state = "OPEN"
}

func IsDoorOpen(ptrDoor *Door) bool{
	PrintStr("is the Door opened ?")
	if ptrDoor.state == "OPEN"{
		return true
	}else {
		return false
	}
}

func IsDoorClose(ptrDoor *Door) bool {
	PrintStr("is the Door closed ?")
	if  ptrDoor.state == "CLOSE" {
		return true
	}else {
		return false
	}
}

func main() {
	door := &Door{}

	OpenDoor(door)
	if IsDoorClose(door) {
		OpenDoor(door)
	}
	if IsDoorOpen(door) {
		CloseDoor(door)
	}
	if door.state == "OPEN" {
		CloseDoor(door)
	}
}
