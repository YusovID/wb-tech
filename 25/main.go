package main

import (
	"fmt"
	"time"
)

var timeToSleep = 1 // секунд

func main() {
	for i := range 10 {
		sleep(timeToSleep)
		fmt.Printf("i: %d, current time: %v\n", i, time.Now().Format("2006-01-02 15:04:05"))
	}
}

func sleep(duration int) {
	ticker := time.NewTicker(time.Duration(duration) * time.Second)
	defer ticker.Stop()

	// блокируемся, пока в тикере не произойдет "тик"
	<-ticker.C
}
