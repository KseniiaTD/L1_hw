package main

// Разработать программу, которая будет последовательно отправлять значения в канал,
// а с другой стороны канала — читать. По истечению N секунд программа должна завершаться.

import (
	"flag"
	"fmt"
	"time"
)

func writer(ch chan<- int, secNum int) { //функция записи в канал
	timeoutCh := time.After(time.Second * time.Duration(secNum)) //отсчет времени, при завершении времени передается значение остановки в канал таймера
	for x := 0; ; x++ {                                          // бесконечный цикл для передачи данных
		select {
		case <-timeoutCh: // если время вышло закрываем канал и выходим из функции
			close(ch)
			fmt.Println("writer closed")
			return
		default: // передача данных в канал
			ch <- x
		}
	}
}

func reader(ch <-chan int) { // функция чтения из канала
	for i := range ch { //читаем данные из канала пока он не будет закрыт
		fmt.Printf("Read %d\n", i)
	}
	fmt.Println("reader closed")
}

func main() {
	ch := make(chan int)                              //создаем канал для передачи данных
	seconds := flag.Int("s", 10, "Timmer in seconds") //при вызове программы указываем время работы программы
	flag.Parse()

	go writer(ch, *seconds) // передаем канал и время работы в горутину
	reader(ch)              // чтение из канала
}
