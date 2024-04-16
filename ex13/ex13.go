package main

// Поменять местами два числа без создания временной переменной.

import "fmt"

func swap1(i int, j int) { // меняем числа через множественное присваивание
	fmt.Println("before i=", i, "j=", j)
	i, j = j, i
	fmt.Println("after i=", i, "j=", j)
}

func swap2(i int, j int) { // меняем числа через сумму и разности
	fmt.Println("before i=", i, "j=", j)
	i = i - j // вычитаем из первого числа второе 15 и 11 -> 4
	j = i + j // разность чисел прибавляем к 2 числу 4 + 11 -> 15 = j
	i = j - i // из получившегося значения вычитаем  разность чисел 15 - 4 -> 11 = i
	fmt.Println("after i=", i, "j=", j)
}

func swap3(i int, j int) { //меняем числа через xor
	fmt.Println("before i=", i, "j=", j)
	i ^= j // выполняем xor для i и j: 101 и 011 ->  110
	j ^= i // 011 и 110 -> 101 = j
	i ^= j // 110 и 101 -> 011 = i
	fmt.Println("after i=", i, "j=", j)
}

func main() {
	i, j := 5, 13 // задаем два числа
	swap1(i, j)
	swap2(i, j)
	swap3(i, j)
}
