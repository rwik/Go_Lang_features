package main

import "fmt"

func main() {

	fmt.Print(sum(10, 20, 34))

}

func sum(a ...int) int {

	total := 0

	for _, num := range a {
		total += num
	}

	return total
}
