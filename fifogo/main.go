package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"syscall"
	"time"
)

func main() {
	fifoPath := "/tmp/myfifo"

	// Проверяем, существует ли FIFO, если нет, создаем его
	if _, err := os.Stat(fifoPath); os.IsNotExist(err) {
		err := syscall.Mkfifo(fifoPath, 0666)
		if err != nil {
			log.Fatalf("Ошибка создания FIFO: %v\n", err)
		}
		fmt.Println("FIFO создано:", fifoPath)
	} else {
		fmt.Println("FIFO уже существует:", fifoPath)
	}

	// Запускаем чтение и запись в отдельных горутинах
	go writeToFIFO(fifoPath)
	go readFromFIFO(fifoPath)

	// Ожидаем завершения операций
	time.Sleep(10 * time.Second) // Даем время для операций чтения/записи
}

// Функция для записи в FIFO
func writeToFIFO(fifoPath string) {
	// Открываем FIFO для записи
	fifo, err := os.OpenFile(fifoPath, os.O_WRONLY, 0666)
	if err != nil {
		log.Fatalf("Ошибка открытия FIFO для записи: %v\n", err)
	}
	defer fifo.Close()

	// Записываем сообщения в FIFO
	messages := []string{"Привет из Go!", "Сообщение 2", "Сообщение 3"}
	for _, msg := range messages {
		fmt.Printf("Отправка сообщения: %s\n", msg)
		_, err := fifo.Write([]byte(msg + "\n"))
		if err != nil {
			log.Fatalf("Ошибка записи в FIFO: %v\n", err)
		}
		time.Sleep(1 * time.Second) // Задержка перед отправкой следующего сообщения
	}
}

// Функция для чтения из FIFO
func readFromFIFO(fifoPath string) {
	// Открываем FIFO для чтения
	fifo, err := os.OpenFile(fifoPath, os.O_RDONLY, 0666)
	if err != nil {
		log.Fatalf("Ошибка открытия FIFO для чтения: %v\n", err)
	}
	defer fifo.Close()

	reader := bufio.NewReader(fifo)
	for {
		// Чтение сообщений
		msg, err := reader.ReadString('\n')
		if err != nil {
			log.Printf("Ошибка чтения из FIFO: %v\n", err)
			break
		}
		fmt.Printf("Получено сообщение: %s", msg)
	}
}
