// Package worker defines the Worker struct and related functions.
// A Worker is responsible for processing data.
package worker

import (
	"bufio"
	"fmt"
	"level-1/3/message"
	"log"
	"os"
	"strconv"
)

// Worker represents a processing unit that can handle tasks.
type Worker struct {
	// Id is the unique identifier for the worker.
	Id int
	// ProcessedDataCount tracks the number of data items this worker has processed.
	ProcessedDataCount int
}

// New creates and initializes a slice of workers.
// It prompts the user to input the desired number of workers.
func New() []*Worker {
	// Get the number of workers from user input.
	workersNum := mustInputWorkersNum()

	// Pre-allocate a slice for the workers with the specified capacity.
	workers := make([]*Worker, 0, workersNum)

	// Create and append new Worker instances to the slice.
	for i := range workersNum {
		workers = append(workers, &Worker{Id: i + 1})
	}

	return workers
}

// mustInputWorkersNum prompts the user to enter the number of workers
// and handles the input parsing. It will terminate the program on error.
func mustInputWorkersNum() int {
	// Create a new scanner to read from standard input.
	scanner := bufio.NewScanner(os.Stdin)

	// Inform the user about the recommended maximum number of workers.
	fmt.Printf("WARNING: Recomended max number of workers is %d!\n", message.MaxMessages)
	fmt.Print("Input number of workers: ")
	// Scan the user's input.
	scanner.Scan()

	// Convert the input text to an integer.
	workersNum, err := strconv.Atoi(scanner.Text())
	if err != nil {
		// If conversion fails, log the error and exit.
		log.Fatalf("failed to convert text to int: %v", err)
	}

	return workersNum
}
