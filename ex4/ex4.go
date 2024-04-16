package main

// Реализовать постоянную запись данных в канал (главный поток). Реализовать набор из N воркеров,
// которые читают произвольные данные из канала и выводят в stdout. Необходима возможность выбора количества
//  воркеров при старте.

// Программа должна завершаться по нажатию Ctrl+C. Выбрать и обосновать способ завершения работы
// всех воркеров.

import (
	"flag"
	"fmt"
	"os"
	"os/signal"
	"sync"
	"syscall"
)

func writer(out chan<- int, chStop <-chan os.Signal) { // функция записи данных в канал
	for x := 0; ; x++ { // бесконечныы2й цикл для записи данных в канал
		select {
		case <-chStop: // если произошло нажатие Ctrl+C, закрываем канал для данных и выходим из функции
			close(out)
			return
		default: // бесконечно передаем данные в канал
			out <- x
		}
	}
}

func reader(in <-chan int, w int, wg *sync.WaitGroup) { // функция чтения из канала
	defer wg.Done()     // по завершению функции уменьшаем список ожжидаемых горутин группы
	for i := range in { // читаем из канала и выводим результат в stdout
		fmt.Printf("worker: %d, number: %d\n", w, i)
	}
	fmt.Printf("worker %d is stopped\n", w) //выводим информацию о закрытии канала после его закрытия в writer
}

func main() {
	workerNum := flag.Int("n", 3, "number of workers") //Указываем количество воркеров при запуске
	flag.Parse()                                       // Парсим флаги командной строки

	ch := make(chan int, *workerNum)                     // создаем буферизированный канал для передачи данных
	chStop := make(chan os.Signal, 1)                    //создаем канал для передачи сигнала остановки
	signal.Notify(chStop, os.Interrupt, syscall.SIGTERM) //перехват сигнала Ctrl+C и запись в канал сигнала

	fmt.Println("workernum", *workerNum)

	var wg sync.WaitGroup              // создаем группу для горутин
	for w := 1; w <= *workerNum; w++ { // создаем набор воркеров
		wg.Add(1)             //увеличиваем счетчик горутин на 1, после завершения горутины он будет уменьшаться
		go reader(ch, w, &wg) //вызов горутины
	}

	writer(ch, chStop) // передача каналов в функцию записи
	wg.Wait()          //ожидаем выполнение всех горутин

	fmt.Println("Graceful shutdown")
}
