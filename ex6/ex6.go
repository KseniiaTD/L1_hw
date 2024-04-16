package main

// Реализовать все возможные способы остановки выполнения горутины.

import (
	"context"
	"fmt"
	"os"
	"time"
)

func case1(sec int) { // бескончно работает, пока время не истечет
	for {
		select {
		case <-time.After(time.Second * time.Duration(sec)): //по завершению времени передастся сигнал
			fmt.Println("Stop by timeout")
			return //завершим функцию при срабатывании таймера
		default:
			//do nothing
		}
	}
}

func case2(context context.Context) { //ункция завершится по сигналу отмены из контекста
	for {
		select {
		case <-context.Done(): // если передали сигнал отмены, завершаем выполнение
			fmt.Println("Stop by context")
			return
		default:
			//do nothing
		}
	}

}

func case3(done <-chan int) { // функция завершится при передаче числа из канала
	for i := range done { // читаем канал пока не встретится 20
		fmt.Println("case3", i)
		if i == 20 {
			fmt.Println("Stop by value from channel")
			return
		}
	}
}

func case4(closed <-chan int) { //функция завершится при закрытии канала
	for i := range closed {
		fmt.Println("case4", i)
	}
	fmt.Println("Stop by closing channel")
}

func writer(out chan<- int, chStop <-chan os.Signal) { //from ex4
	//пример завершения функции по нажатию клавиши
	//здесь только часть кода, весь код в ex4
	for x := 0; ; x++ {
		select {
		case <-chStop:
			close(out)
			return
		default:
			out <- x
		}
	}
}

func main() {
	go case1(5) // горутина работающая по таймеру

	context, cancel := context.WithCancel(context.Background()) //создаем контекст с отменой по явному сигналу
	go case2(context)                                           // передаем в горутину контекст
	time.Sleep(5 * time.Second)
	cancel() // после 5 секунд ожидания вызываем отмену операции, горутина завершается

	done := make(chan int)   // создаем канал для case3
	closed := make(chan int) // создаем канал для case4

	go case3(done)   //остановка горутины при передаче определенного значения
	go case4(closed) //завершение горутины после закрытия канала
	for i := 0; i < 20; i++ {
		done <- i + 1 // при передаче значения 20 горутина case3 завершится
		closed <- i
	}
	close(closed) //закрываем канал, горутина case4 завершится

	time.Sleep(3 * time.Second) // нужно, чтобы горутина 4 успела вывести сообщение
}
