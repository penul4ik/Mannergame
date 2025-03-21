package main

import (
	"fmt"
	"io"
	"net"
	"net/http"
	"os"
	"time"
)

func workerHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Worker: Accept the request from client")
	time.Sleep(500 * time.Millisecond)

	body, err := io.ReadAll(r.Body)
	if err != nil {
		fmt.Println("Worker: Error while read requset's body")
		http.Error(w, "Server error", http.StatusInternalServerError)
		return
	}

	defer r.Body.Close()

	fmt.Println("Worker: Data:", string(body))

	fmt.Println("Worker: Read file from disk...")
	time.Sleep(1 * time.Second)

	fileContent, err := os.ReadFile("response.txt")
	if err != nil {
		fmt.Println("Worker: Read file error")
		http.Error(w, "Server error", http.StatusInternalServerError)
		return
	}

	fmt.Println("Worker: Send HTTP-response to client")
	w.Write(fileContent)
}

func main() {
	listener, err := net.Listen("tcp", ":8080")
	if err != nil {
		fmt.Println("Starting server error:", err)
		return
	}

	fmt.Println("Server started on the 8080 port")
	go func() {
		http.HandleFunc("/", workerHandler)
		http.Serve(listener, nil)
	}()

	// Симуляция клиента
	time.Sleep(1 * time.Second) // Ждем, пока сервер запустится

	fmt.Println("Клиент: Отправляет HTTP-запрос...")
	time.Sleep(500 * time.Millisecond)

	resp, err := http.Get("http://localhost:8080")
	if err != nil {
		fmt.Println("Клиент: Ошибка при отправке запроса")
		return
	}
	defer resp.Body.Close()

	fmt.Println("Клиент: Получает ответ...")
	responseData, _ := io.ReadAll(resp.Body)
	fmt.Println("Клиент: Ответ сервера:", string(responseData))
}
