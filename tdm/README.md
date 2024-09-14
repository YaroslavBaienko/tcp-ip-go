# Go TCP Server with Time Division Multiplexing (TDM)

This is a simple Go-based TCP server that simulates **Time Division Multiplexing (TDM)**. The server handles multiple client connections and processes each one in a separate time slot. The time slots ensure that clients access a shared resource (in this case, a log file) sequentially and in a controlled manner.

## Features

- **Time Division Multiplexing (TDM)**: Clients are assigned a time slot and wait for their turn to write data to the shared file.
- **Concurrency**: Each client connection is processed in a separate goroutine.
- **File synchronization**: Access to the shared log file is protected using a mutex to prevent concurrent writes.
- **TCP communication**: Clients can connect to the server via TCP and receive confirmation when their request is processed.

## How It Works

1. When a client connects, the server assigns it a time slot.
2. Each client waits for its designated time slot to be processed (simulating TDM).
3. Once the slot is reached, the client writes a message to a shared log file (`server_log.txt`).
4. The server ensures that only one client can write to the file at a time using a mutex.
5. After processing, the server notifies the client that its request has been handled.

## Requirements

- Go 1.16 or higher.
