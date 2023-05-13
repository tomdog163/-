package day01

import (
	"fmt"
	"math"
)

func Example01() {
	var a = "initial"
	var b, c int = 1, 2
	var d = true
	var e float64

	f := float32(e)
	g := a + "foo"

	fmt.Println(a, b, c, d, e, f)
	fmt.Println(g)

	const s string = "constant"
	const h = 50000000
	const i = 3e20 / h
	fmt.Println(s, h, i, math.Sin(h), math.Sin(i))
}
