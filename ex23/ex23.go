package main

// Удалить i-ый элемент из слайса.

import (
	"flag"
	"fmt"
)

func main() {
	elem := flag.Int("e", 0, "Position of element for delete operation") // передаем номер удаляемого из слайса элемента
	flag.Parse()                                                         // парсим флаги
	n := 10
	deletedElem := *elem                     // присваиваем значение флага
	if n <= deletedElem || deletedElem < 0 { // проверка на выход за границы слайса
		fmt.Println("index out of range")
		return
	}

	slice := make([]int, n) // создаем слайс

	for i := 0; i < n; i++ {
		slice[i] = i + 1
	}

	sliceLeft := slice[0:deletedElem]        // создаем слайс как срез значений до удаляемого элемента
	sliceRight := slice[deletedElem+1:]      // создаем слайс как срез значений после удаляемого элемента
	slice = append(sliceLeft, sliceRight...) // объединяем слайсы

	fmt.Println(slice) // выводим результирующий слайс
}
