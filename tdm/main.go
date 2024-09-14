package main

import (
	"fmt"
	"net"
	"os"
	"sync"
	"time"
)

// Time slot duration for each client (simulates TDM slots)
const slotDuration = 2 * time.Second

// Mutex for controlling access to the file
var fileMutex sync.Mutex

// Function to handle each client's request
func handleConnection(conn net.Conn, id int, slots chan int) {
	defer conn.Close()

	// Receive a time slot for processing
	slot := <-slots

	// Wait until it's the client's turn (simulating time slot assignment)
	fmt.Printf("Client %d is waiting for its turn\n", id)
	time.Sleep(time.Duration(slot) * slotDuration)

	// Lock and write data to the file
	fileMutex.Lock()
	file, err := os.OpenFile("server_log.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Printf("Error opening file: %v\n", err)
		fileMutex.Unlock()
		return
	}
	defer file.Close()

	// Write the client's message to the file
	message := fmt.Sprintf("Client %d writes data to the file at %s\n", id, time.Now().Format(time.RFC1123))
	file.WriteString(message)
	fmt.Printf("Client %d has written data: %s", id, message)

	// Unlock the file
	fileMutex.Unlock()

	// Notify the client that their request has been processed
	conn.Write([]byte("Your request has been processed!\n"))
}

func main() {
	// Create a channel for time slots
	slots := make(chan int, 10)

	// Initialize the slots
	go func() {
		slot := 0
		for {
			slots <- slot
			slot = (slot + 1) % 10 // Total available slots, wrap around when max is reached
		}
	}()

	// Start the server
	listener, err := net.Listen("tcp", ":8080")
	if err != nil {
		fmt.Printf("Error starting server: %v\n", err)
		os.Exit(1)
	}
	defer listener.Close()

	fmt.Println("Server started and waiting for connections on port 8080...")

	// Handle incoming connections
	clientID := 1
	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Printf("Error accepting client connection: %v\n", err)
			continue
		}

		// Start a new goroutine to handle the client's request
		go handleConnection(conn, clientID, slots)
		clientID++
	}
}
