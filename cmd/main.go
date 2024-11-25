package main

import (
	"fmt"
	"os"
)

func main() {
	// Открываем файл для логов
	logFile, err := os.OpenFile("app.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Println("Ошибка открытия файла:", err)
		return
	}
	defer logFile.Close() // Закрываем файл при завершении работы программы

	// Записываем несколько сообщений в лог
	for i := 0; i < 5; i++ {
		logMessage := fmt.Sprintf("Запись %d в лог\n", i+1)
		_, err := logFile.WriteString(logMessage)
		if err != nil {
			fmt.Println("Ошибка записи в лог:", err)
			return
		}
	}
	fmt.Println("Логирование завершено")
}
