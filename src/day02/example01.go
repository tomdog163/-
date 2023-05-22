package day02

import (
	"fmt"
	"math/rand"
	"strconv"
)

// Example01 猜谜游戏
func Example01() {
	maxNum := 100
	secretNumber := rand.Intn(maxNum)
	//fmt.Println("The secret number is", secretNumber)

	fmt.Println("please input your guess")

	for {
		var input string
		fmt.Scan(&input)

		// 转换成数字
		guess, err := strconv.Atoi(input)

		if err != nil {
			fmt.Println("Invalid input. Please enter on integer value")
			return
		}
		fmt.Println("You guess is", guess)

		if guess > secretNumber {
			fmt.Println("You guess is bigger than secret. Please try again")
		} else if guess < secretNumber {
			fmt.Println("You guess is smaller than secret. Please try again")
		} else {
			fmt.Println("Correct, you legend")
			break
		}
	}
}
