package main

import (
	"fmt"
	"math/rand/v2"
	"time"
)

const (
	maxNum = 1000 // максимальное число, которое может быть сгенерировано
	TTL    = 3    // время жизни программы в секундах
)

func main() {
	out := generateRandom()

	// используем оба варианта для демонстрации работы в разных случаях
	afterChan := time.After(TTL * time.Second)
	timer := time.NewTimer(TTL * time.Second)

	for {
		select {
		case <-afterChan:
			fmt.Println("Отработал time.After()!")
			return

		case <-timer.C:
			fmt.Println("Отработал NewTimer()!")
			return

		case num := <-out:
			fmt.Printf("got new number: %d\n", num)
		}
	}
}

// generateRandom генерирует случайное число в диапазоне от 0 до 1000 каждые 100 миллисекунд
// и записывает их в канал
func generateRandom() chan int {
	out := make(chan int)

	go func() {
		for {
			randNum := rand.IntN(maxNum + 1) // +1, потому что результат IntN(n) = [0, n)
			out <- randNum
			time.Sleep(100 * time.Millisecond)
		}
	}()

	return out
}
