package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	t := time.Now().UnixNano()
	rand.Seed(t)
	// xは1-6の間の値になる
	num := rand.Intn(6) + 1
	switch num {
	case 1:
		fmt.Println("今日の運勢は凶です")
	case 2, 3:
		fmt.Println("今日の運勢は吉です")
	case 4, 5:
		fmt.Println("今日の運勢は中吉です")
	case 6:
		fmt.Println("今日の運勢は大吉です")
	}
}
