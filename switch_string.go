package main

import "fmt"

func main() {

	key := "hi"
	switch key {
	case "hi":
		fmt.Printf("case 1")
	default:
		fmt.Printf("case default")
	}
}
