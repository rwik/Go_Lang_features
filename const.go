package main

import "fmt"

func main() {
	const pi = 3.14
	//pi = pi / 2 //error

	const (
		a = 34
		b
		c
		d = 10
		e
		f
		g
	)
	fmt.Println(a, b, c, d, e, f, g)
	//34 34 34 10 10 10 10

}
