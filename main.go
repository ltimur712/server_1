package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/ok", okHandler)   // Обработчик для страницы /ok
	http.HandleFunc("/info", infoHandler) // Обработчик для страницы /info

	// Запуск веб-сервера на порту 8080
	fmt.Println("Сервер запущен на http://localhost:8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		fmt.Println("Ошибка запуска сервера:", err)
	}
}

// Обработчик для страницы "OK"
func okHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Status: OK"))
}

// Обработчик для страницы "Info"
func infoHandler(w http.ResponseWriter, r *http.Request) {
	studentInfo := "ФИО: Иванов Иван Иванович\nГруппа: 123-45"
	w.WriteHeader(http.StatusOK) 
	w.Write([]byte(studentInfo))
}