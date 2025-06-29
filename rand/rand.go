package rand

import (
	"fmt"
	"time"
)

func simpleRandom(seed int64) int {
	// Простейший хэш на основе побитовых операций и арифметики
	seed ^= (seed << 13)
	seed ^= (seed >> 7)
	seed ^= (seed << 17)
	return int(seed & 0x7FFFFFFF)
}

func Random() int {
	// Используем время как источник "энтропии"
	seed := time.Now().UnixNano()
	random := simpleRandom(seed) % 100 // от 0 до 99
	fmt.Println("Случайное число:", random)
	return random
}
