package rand

import (
	"fmt"

	"github.com/k0kubun/pp"
)

func test() {
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
	Random()

	for i := 0; i < 100; i++ {
		fmt.Println(i, "-", factorial(int64(i)))
	}

	fmt.Println(factorial(10))
}

func factorial(x int64) int64 {
	if x == 0 {
		return 1
	}

	return x * factorial(x-1)
}
