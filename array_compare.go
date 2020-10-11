package main

import (
	"fmt"
	"reflect"
)

func main() {
	a := []int{1, 2, 3}
	b := []int{1, 2, 3}

	fmt.Println(reflect.DeepEqual(a, b))

	for i, x := range a {
		if x != b[i] {
			fmt.Print("false")
		}
	}

}
