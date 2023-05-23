package day03

import "fmt"

// Example02
func Example02() {
	CalSquare()
}

func CalSquare() {
	src := make(chan int)
	dest := make(chan int, 3)
	// A协程发送 0-9 数字
	go func() {
		defer close(src)
		for i := 0; i < 10; i++ {
			src <- i
		}
	}()
	// B协程计算输入数字的平方
	go func() {
		defer close(dest)
		for i := range src {
			dest <- i * i
		}
	}()
	// 主协程输出最后的平方数
	for i := range dest {
		// 复杂操作
		fmt.Println(i)
	}
}
