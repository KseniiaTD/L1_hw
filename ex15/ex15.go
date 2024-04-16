package main

// К каким негативным последствиям может привести данный фрагмент кода, и как это исправить?
// Приведите корректный пример реализации.

// var justString string
// func someFunc() {
//   v := createHugeString(1 << 10)
//   justString = v[:100] -- если символ занимает больше 1 байта(например, русские буквы занимают 2 байта, иероглифы - 3), может произойти "разрыв" символа
// }

// func main() {
//   someFunc()
// }

import (
	"fmt"
	"strings"
)

func createHugeString(i int) string {
	return strings.Repeat("в", i) // например, задаем строку из "в" в i символов
}

var justString string

func someFunc() {
	v := createHugeString(1 << 10)
	//justString = v[:100] //некорректная работа для символов больше 1 байта (русские буквы, иероглифы и т.д.)
	arr := []rune(v)               // создаем слайс из рун, руны могут содержать больше 1 байта
	justString = string(arr[:100]) // преобразуем срез из 100 рун в строку
	fmt.Println(justString)
}

func main() {
	someFunc()
}
