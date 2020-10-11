package main

import "fmt"

func main() {
	daysMap := make(map[string]int)

	fmt.Print(daysMap)
	daysMap["sat"] = 10
	fmt.Println(daysMap)

}
