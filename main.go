package main

import (
	"fmt"
	"net"
	"net/http"
	"strings"
	"time"
)

func main() {
	http.HandleFunc("/ok", okHandler)       // Обработчик для страницы /ok
	http.HandleFunc("/info", infoHandler)   // Обработчик для страницы /info
	http.HandleFunc("/status", statusHandler) // Обработчик для страницы /status

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
	studentInfo := "ФИО: Лобин Тимур Сергеевич\nГруппа: 712"
	w.WriteHeader(http.StatusOK) 
	w.Write([]byte(studentInfo)) 
}

// Обработчик для страницы "Status"
func statusHandler(w http.ResponseWriter, r *http.Request) {
	// Получение IP-адреса пользователя
	ip := getMaskedIP(r)

	// Получение текущего времени
	currentTime := time.Now().Format("2006-01-02 3:4:5 pm")

	// Информация для вывода
	response := fmt.Sprintf("IP-адрес: %s\nФИО: Лобин Тимур Сергеевич\nТекущее время: %s", ip, currentTime)

	w.WriteHeader(http.StatusOK)
	w.Write([]byte(response))    
}

func getMaskedIP(r *http.Request) string {
	// Получаем IP-адрес из заголовка
	ip, _, err := net.SplitHostPort(r.RemoteAddr)
	if err != nil {
		return "Неизвестный IP"
	}

	// Разбиваем IP на октеты
	octets := strings.Split(ip, ".")
	if len(octets) == 4 {
		octets[1] = "XXX" 
		octets[2] = "XXX"
		return strings.Join(octets, ".")
	}

	return ip 
}
