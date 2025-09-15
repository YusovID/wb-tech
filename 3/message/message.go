// Package message defines the Message struct and functions for generating and processing messages.
package message

import (
	"fmt"
	"math/rand/v2"
	"sync"
	"time"
)

// MaxMessages defines the maximum number of messages to be retrieved in a single batch.
const MaxMessages = 10

// Message represents a data unit to be processed.
type Message struct {
	// id is the unique identifier for the message.
	id int
	// text contains the content of the message.
	text string
}

// Generate creates a channel and starts a goroutine to continuously produce messages.
// It returns the channel, which can be used to receive the generated messages.
func Generate() chan *Message {
	// Create an unbuffered channel for Messages.
	messagesChan := make(chan *Message)

	// Start a new goroutine for message generation.
	go func() {
		i := 0
		// Loop indefinitely to create and send messages.
		for {
			message := &Message{
				id:   i + 1,
				text: "", // Text is empty for this example.
			}
			// Send the newly created message to the channel.
			messagesChan <- message
			i++
		}
	}()

	return messagesChan
}

// Get retrieves a random number of messages from the provided channel.
// It fetches up to MaxMessages in a single call.
func Get(messagesChan chan *Message) []*Message {
	// Pre-allocate a slice with a capacity of MaxMessages.
	messages := make([]*Message, 0, MaxMessages)

	wg := &sync.WaitGroup{}

	// Use wg.Go which is a convenient way to use WaitGroup
	wg.Go(
		func() {
			// Determine a random number of messages to get for this batch.
			messagesCnt := rand.IntN(MaxMessages + 1)

			// Read the specified number of messages from the channel.
			for range messagesCnt {
				message := <-messagesChan
				messages = append(messages, message)
			}
		})

	// Wait for the goroutine to finish fetching messages.
	wg.Wait()

	return messages
}

// Process is a handler function that simulates the processing of a single message.
// It prints a message to stdout indicating which worker processed which message.
func Process(workerId int, message *Message) {
	// Simulate work by sleeping for a short duration.
	time.Sleep(100 * time.Millisecond)
	fmt.Printf("Worker %d got message %d\n", workerId, message.id)
}
