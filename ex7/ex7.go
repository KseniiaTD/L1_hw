package main

// Реализовать конкурентную запись данных в map.

import (
	"fmt"
	"math/rand"
	"sync"
)

func main() {
	m := make(map[string]int) //создаем пустую мапу
	mu := sync.Mutex{}        //создаем мьютекс для синхронизации записи в мапу

	var letterRunes = []rune("ABCDEFGHIJKLMNOPQRSTUVWXYZ") // создаем слаайс рун из строки, будем подсчитывать кол-во повторений букв

	var wg sync.WaitGroup     //задаем группу для горутин
	for i := 0; i < 10; i++ { // запускаем 10 горутин
		wg.Add(1)
		go func(i int) {
			defer wg.Done()              // по завершению горутины уменьшаем счетчик группы на 1
			for j := 0; j < 10000; j++ { //10000 раз получаем букву и считаем кол-во ее повторений
				key := string(letterRunes[rand.Intn(len(letterRunes))]) //рандомное получение буквы из слайса рун
				mu.Lock()                                               //блокируем мапу для синхронизации
				m[key]++                                                //увеличиваем повторение буквы на 1
				mu.Unlock()                                             //снимаем блокировку
			}
		}(i)
	}

	wg.Wait() // ждем завершения всех горутин
	fmt.Println(m)
}
