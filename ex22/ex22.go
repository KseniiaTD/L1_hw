package main

// Разработать программу, которая перемножает, делит, складывает, вычитает
// две числовых переменных a,b, значение которых > 2^20.

import (
	"fmt"
	"math/big"
)

func main() { // для расчета больших чисел используем библиотеку big
	a := big.NewInt(2000000100777) // задаем 2 числа
	b := big.NewInt(3000000002)

	mul := new(big.Int) // перемножение больших чисел
	mul.Mul(a, b)
	fmt.Println("Mul:", mul)

	sum := new(big.Int) // сложение больших чисел
	sum.Add(a, b)
	fmt.Println("Sum:", sum)

	div := new(big.Int) // деление больших чисел
	div.Div(a, b)
	fmt.Println("Div:", div)

	sub := new(big.Int) // вычитание больших чисел
	sub.Sub(a, b)
	fmt.Println("Sub:", sub)
}
