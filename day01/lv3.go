package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	maxNum := 100
	rand.Seed(time.Now().UnixNano())
	secretNumber := rand.Intn(maxNum)
	fmt.Println("The secret number is ", secretNumber)

	fmt.Println("Please input your guess")
	var guess int

	for  true {
		// 输入我们猜的数字
		_, err := fmt.Scanf("%d", &guess)
		// Go语言中处理错误的方法
		if err != nil {
			fmt.Println("Invalid input. Please enter an integer value")
			return
		}
		if guess == secretNumber {
			fmt.Printf("Congratulations！You are right！Answer is %d",guess)
			return
		} else if guess >secretNumber{
			fmt.Println("Sorry.It's too big.")
		}else if guess <secretNumber{
			fmt.Println("Sorry.It's too small.")
		}
	}
}