Go FIFO Communication Example
Overview

This project demonstrates how to use FIFO (First In, First Out) for inter-process communication in Go. The code creates a FIFO file in /tmp/myfifo and utilizes two goroutines for writing and reading messages through this file. The writer sends a series of messages to the FIFO, and the reader reads and displays them.
Prerequisites

Before running the program, make sure you have:

    Go installed on your machine.
    A Unix-like system that supports FIFO (e.g., Linux or macOS).

How It Works

    FIFO Creation:
        The program checks if a FIFO file already exists at /tmp/myfifo. If not, it creates one using syscall.Mkfifo.

    Writing to FIFO:
        The writeToFIFO function writes a series of messages to the FIFO file with a 1-second delay between each message. This is done in a separate goroutine.

    Reading from FIFO:
        The readFromFIFO function reads messages from the FIFO in another goroutine, using a buffered reader to capture each message line by line.

    Concurrency:
        Both writing and reading are executed concurrently via goroutines. A time.Sleep(10 * time.Second) ensures the program runs long enough to complete all operations.

Code Breakdown

    os.Stat() and syscall.Mkfifo: Used to check the existence of the FIFO file and create it if it doesn’t exist.

    os.OpenFile(): Opens the FIFO file for reading or writing depending on the mode specified.

    bufio.NewReader(): Wraps the FIFO file in a buffered reader for efficient reading of lines.

    Goroutines: The program uses Go’s concurrency model to simultaneously write and read messages using goroutines.

Key Concepts

    FIFO: A special file type that follows the First In, First Out principle for reading and writing.
    syscall.Mkfifo: A Go system call used to create FIFO files.
    Concurrency: Writing and reading from the FIFO are handled in separate goroutines to demonstrate Go’s concurrency capabilities.

Error Handling

    If the FIFO creation or access fails, the program will log a fatal error using log.Fatalf.
    Errors in reading or writing are gracefully handled, ensuring the program continues its operation.

Customization

    FIFO Path: You can change the FIFO path by modifying the fifoPath variable.
    Messages: Modify the messages array in the writeToFIFO function to customize the messages sent through the FIFO.

Further Learning

    Goroutines and Channels: To deepen your understanding of concurrency in Go.
    Inter-Process Communication (IPC): Explore different IPC mechanisms like pipes, message queues, and shared memory.

Author

Yaroslav Baienko - DevOps and SysOps Engineer. You can find more about my projects on GitHub​.
