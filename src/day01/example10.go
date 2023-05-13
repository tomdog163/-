package day01

import "fmt"

func add2(n int) {
	n += 2
}

func add2ptr(n *int) {
	*n += 2
}

func Example10() {
	n := 5
	add2(n)
	fmt.Println(n)
	add2ptr(&n)
	fmt.Println(n)
}
