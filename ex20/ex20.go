package main

// Разработать программу, которая переворачивает слова в строке.
// Пример: «snow dog sun — sun dog snow».

import (
	"fmt"
	"strings"
)

func main() {
	str := "Клара у Карлы украла кораллы" // задаем строку

	words := strings.Split(str, " ") // разбиваем строку на слайс слов

	for i := range words { // проходим в цикле слайс слов
		if i > len(words)/2 { // если прошли больше половины, выходим из цикла
			break
		}
		words[i], words[len(words)-i-1] = words[len(words)-i-1], words[i] // меняем местами i и n-i слова
	}

	strNew := strings.Join(words, " ") // собираем слова в строку

	fmt.Println(str, "->", strNew) // выводим результат
}
