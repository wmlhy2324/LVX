package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())

	number := rand.Intn(100) + 1
	var left int = 0
	var right int = 100

	for left <= right {
		var mid int = (left + right) / 2
		if mid < number {
			left = mid + 1

		} else if mid > number {
			right = mid - 1
		} else {
			fmt.Println("找到啦这个是", mid)
			break

		}

	}

}
