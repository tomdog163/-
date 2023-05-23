package day03

import (
	"fmt"
	"sync"
	"time"
)

var (
	x    int64
	lock sync.Mutex
)

// Example03 并发安全问题
func Example03() {
	Add()
}

func addWithLock() {
	for i := 0; i < 2000; i++ {
		lock.Lock()
		x += 1
		lock.Unlock()
	}
}
func addWithoutLock() {
	for i := 0; i < 2000; i++ {
		x += 1
	}
}
func Add() {
	x = 0
	for i := 0; i < 5; i++ {
		go addWithoutLock()
	}
	time.Sleep(time.Second)
	fmt.Println("WithoutLock:", x)
	x = 0
	for i := 0; i < 5; i++ {
		go addWithLock()
	}
	time.Sleep(time.Second)
	fmt.Println("WithLock:", x)
}
