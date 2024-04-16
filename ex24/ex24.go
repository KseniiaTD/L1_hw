package main

// Разработать программу нахождения расстояния между двумя точками, которые представлены в виде структуры Point
// с инкапсулированными параметрами x,y и конструктором.

import (
	"fmt"
	"math"
)

type Point struct { // создаем структуру Точка
	x int // координата x
	y int // координата y
}

func NewPoint(x int, y int) Point { // конструктор Point
	return Point{x: x, y: y}
}

func (p *Point) Distance(otherP Point) float64 { // метод расчета расстояния от текущей точки до любой другой точки
	sqr1 := math.Pow(float64(otherP.x)-float64(p.x), 2) // считаем квадрат разности координат х
	sqr2 := math.Pow(float64(otherP.y)-float64(p.y), 2) // считаем квадрат разности координат y
	return math.Sqrt(sqr1 + sqr2)                       // складываем квадраты
}

func Distance(p1 Point, p2 Point) float64 { // функция рассчета расстояния между точками
	sqr1 := math.Pow(float64(p2.x)-float64(p1.x), 2) // считаем квадрат разности координат х
	sqr2 := math.Pow(float64(p2.y)-float64(p1.y), 2) // считаем квадрат разности координат y
	return math.Sqrt(sqr1 + sqr2)                    // складываем квадраты
}

func main() {
	p1 := NewPoint(5, 1) // задаем две точки
	p2 := NewPoint(12, 13)
	// все три расчета вернут одинаковые значения
	fmt.Println("Distance between points:", Distance(p1, p2)) // расчитываем дистанцию по функции
	fmt.Println("Distance between points:", p1.Distance(p2))  // расчитываем дистанцию по методу
	fmt.Println("Distance between points:", p2.Distance(p1))  // расчитываем дистанцию по методу
}
