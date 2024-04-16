package main

// Разработать программу, которая проверяет, что все символы в строке уникальные (true — если уникальные, false etc). Функция проверки должна быть регистронезависимой.

// Например:
// abcd — true
// abCdefAaf — false
// aabcd — false

import (
	"fmt"
	"strings"
)

func checkUniq(str string) bool { // функция возвращает true, если все символы в строке уникальны
	str = strings.ToUpper(str)          // приводим все символы к верхнему регистру
	mapRunes := make(map[rune]struct{}) // создаем пустое множество рун
	for _, elem := range str {          // в цикле проходим по строке, вытаскивая каждый символ как руну
		_, ok := mapRunes[elem] // если руна уже есть в строке, выходим из функции, возвращая false
		if ok {
			return false
		}
		mapRunes[elem] = struct{}{} // записываем руну в множество
	}
	return true // если в множества пыталичь записать каждый елемент только 1 раз,
	//то строка содержит только уникальные значения, возвращаем true
}

func main() {
	str := "УКabCdefA"          //  задаем строку
	fmt.Println(checkUniq(str)) // проверяем строку на уникальность символов
}
