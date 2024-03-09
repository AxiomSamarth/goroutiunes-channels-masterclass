package main

import "strconv"

// sender sends a messageCount number of messages to the message channel
// and then sends a signal to the done channel to indicate that it has
// finished sending messages.
func sender(messageCount int, message chan string, done chan struct{}) {
	for i := 0; i < messageCount; i++ {
		message <- "Hello World " + strconv.Itoa(i)
	}
	done <- struct{}{}
}

// receiver receives messages from the message channel and prints them
func receiver(message chan string) {
	for msg := range message {
		println(msg)
	}
}

func main() {
	// Create a message channel and a done channel
	message := make(chan string)
	done := make(chan struct{})

	// Start the sender and receiver goroutines
	messageCount := 10

	go sender(messageCount, message, done)
	go receiver(message)

	// Wait for the sender to finish sending messages
	<-done

	// Close the message channel. It is important to close the message channel
	// to indicate that no more messages will be sent on the channel. This is
	// necessary to avoid a deadlock in the receiver goroutine.
	close(message)

	// NOTE:
	// If we do not use the done channel, it will result in the termination of this process
	// before the receiver has a chance to print all the messages. Hence, we use the done
	// channel to wait until the sender has finished sending messages and then we close the
	// message channel and then let the process end.
}
