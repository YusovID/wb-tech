package main

import (
	"fmt"
	"log"

	"github.com/YusovID/wb-tech/level-2/8/internal/ntp"
)

func main() {
	currentTime, err := ntp.GetCurrentTime()
	if err != nil {
		log.Fatalf("ERROR: getCurrentTime failed: %v", err)
	}

	fmt.Printf("current time: %v", currentTime)
}
