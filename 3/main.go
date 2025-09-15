// Package main is the entry point of the application.
// It sets up a pool of workers to process messages concurrently.
package main

import (
	// m is an alias for the message package.
	m "level-1/3/message"
	"level-1/3/pool"
	"os"
	"os/signal"
	"syscall"
)

// IPool defines the interface for a worker pool.
// It outlines the essential operations for managing and interacting with a pool of workers.
type IPool interface {
	// Create initializes the pool by making all workers available.
	Create()
	// Handle submits a message to the pool for processing by an available worker.
	Handle(*m.Message)
	// Wait blocks until all workers have finished their current tasks and are available again.
	Wait()
	// Stats prints the processing statistics for each worker in the pool.
	Stats()
}

// main is the primary function that orchestrates the application's lifecycle.
// It initializes a message channel, creates a worker pool, and handles graceful shutdown.
func main() {
	// Generate starts a goroutine to continuously produce messages and returns a channel to receive them.
	messagesChan := m.Generate()

	// New creates a new worker pool with a specified handler function (m.Process) for processing messages.
	var pool IPool = pool.New(m.Process)

	// sigChan is a channel to receive OS signals for graceful termination.
	sigChan := make(chan os.Signal, 1)
	// Notify directs the specified signals (SIGHUP, SIGINT, SIGTERM, SIGQUIT) to sigChan.
	signal.Notify(sigChan, syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)

	// Start a new goroutine for the main processing loop.
	go func() {
		// This loop continuously fetches and processes batches of messages.
		for {
			// Get retrieves a batch of messages from the message channel.
			var messages []*m.Message = m.Get(messagesChan)

			// Create populates the pool with available workers for the current batch.
			pool.Create()

			// Iterate over the batch of messages and handle each one.
			for _, message := range messages {
				// Handle assigns the message to an available worker for processing.
				pool.Handle(message)
			}

			// Wait ensures all messages in the current batch are processed before starting the next.
			pool.Wait()
		}
	}()

	// Block until a shutdown signal is received on sigChan.
	<-sigChan

	// Print the final statistics after receiving a shutdown signal.
	pool.Stats()
}
