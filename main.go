package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())
	num := rand.Intn(10)
	b := isEven(num)
	fmt.Printf("число %d четное? %v", num, b)
}

func isEven(i int) bool {
	if i%2 == 0 {
		return true
	}
	return false
}
