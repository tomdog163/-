package day03

import (
	"fmt"
	"sync"
)

// Example01 快速打印hello goroutine
func Example01() {
	HelloGoRoutine()
}

func hello(i int) {
	fmt.Println("hello goroutine: " + fmt.Sprint(i))
}

func HelloGoRoutine() {
	var wg sync.WaitGroup
	wg.Add(5)
	for i := 0; i < 5; i++ {
		go func(j int) {
			defer wg.Done()
			hello(j)
		}(i)
	}
	wg.Wait()
}
