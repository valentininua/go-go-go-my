package main

import (
	"myproject/rand"

	"github.com/k0kubun/pp"
)

func main() {
	arr := []int{10, 20, 30, 40, 50}
	arr = append(arr, 60)
	user := map[string]any{
		"name":      "John",
		"age":       "30",
		"city":      "New York",
		"job":       "Developer",
		"salary":    "1000",
		"isMarried": arr,
	}

	pp.Println(user)
	pp.Println(arr)
	rand.Random()
}

func factorial(x uint) uint {
	if x == 0 {
		return 1
	}

	return x * factorial(x-1)
}
