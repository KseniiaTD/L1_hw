package main

// Реализовать собственную функцию sleep.

import (
	"fmt"
	"time"
)

func sleep(t time.Duration) { //реализуем ожидание через канал time.After, в который передаем время задержки
	<-time.After(t) // через заданное время в канал поступит значение для выхода из функции
}

func main() {
	fmt.Println(time.Now()) // печатаем текущее время
	sleep(13 * time.Second) // выполняем sleep на t секунд
	fmt.Println(time.Now()) // печатаем текущее время
}
