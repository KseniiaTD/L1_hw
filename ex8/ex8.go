package main

// Дана переменная int64. Разработать программу которая устанавливает i-й бит в 1 или 0.

import (
	"flag"
	"fmt"
)

func main() {
	num := *flag.Int64("n", 0, "number for bites manipulation") // передаем на входе число
	elemI := *flag.Int64("i", 0, "element i for reverse a bit") // передаем номер изменяемого бита в числе
	flag.Parse()                                                // парсим флаги

	var bitI int64 = 1 << elemI //создаем маску для операции xor (например, 1000 - для изменения 4го бита)
	var res int64 = num ^ bitI  // выполняем операцию xor для изменения i-го бита с помощью маски

	fmt.Println(res)
}
