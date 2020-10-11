package main

import (
	"fmt"
	"strconv"
)

func main() {
	nums := [4]int{1, 2, 3, 4}
	for i := range nums {
		fmt.Printf(strconv.Itoa(i))
	}

}
