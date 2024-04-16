package main

// Дана структура Human (с произвольным набором полей и методов).
// Реализовать встраивание методов в структуре Action от родительской структуры Human (аналог наследования).

import (
	"fmt"
)

type Human struct {
	name  string  //Имя человека
	age   int     //Возраст человека
	maxKm float32 //Максимальная дистанция в км, которую человек может пройти не устав
}

// Получение возраста человека
func (h *Human) getAge() int {
	return h.age
}

// Получение имени человека
func (h *Human) getName() string {
	return h.name
}

// Увеличение возраста человека
func (h *Human) incAge() {
	h.age += 1          //Увеличиваем возраст на год
	h.setMaxKm(h.maxKm) //Обновляем максимальную дистанцию для человека
}

// Установка максимальной дистанции для человека в зависимости от возраста
func (h *Human) setMaxKm(Km float32) {
	if h.age > 10 && h.age < 60 {
		h.maxKm = Km
	} else if h.age < 5 || h.age > 80 {
		h.maxKm = 0
	} else if Km > 2 {
		h.maxKm = 2
	} else {
		h.maxKm = Km
	}
}

// Получение максимальной дистанции для человека
func (h *Human) getMaxKm() float32 {
	return h.maxKm
}

// Структура - наследник
type Walker struct {
	Human              //стуркура Action наследует все поля и методы Human
	distanceKm float32 //Новое поле - пройденная дистанция
}

// Проверка может ли человек пройти еще с учетом уже пройденной дистанции
// bool - закладка на будущее
func (a *Walker) walk(km float32) bool {

	//проверяем может ли человек еще пройти указанное кол-во км
	//если да, увеличиваем пройденную дистанцию
	if a.distanceKm+km <= a.maxKm {
		a.distanceKm += km
		fmt.Printf("Human %s walked %.3f km\n", a.name, km)
		return true
	}
	fmt.Printf("Human %s is tired\n", a.name)
	return false
}

func main() {
	a := Walker{}
	a.age, a.name = 18, "Tim"
	a.setMaxKm(5)
	fmt.Printf("%s's age is %d\n", a.getName(), a.getAge())
	_ = a.walk(3.5)
	_ = a.walk(1.23)
	_ = a.walk(1.4)

	a2 := Walker{}
	a2.age, a2.name = 59, "Bob"
	a2.setMaxKm(50)
	fmt.Printf("%s's age is %d\n", a2.getName(), a2.getAge())
	fmt.Printf("Human %s can go %.2f km\n", a2.getName(), a2.getMaxKm())

	a2.incAge()
	fmt.Printf("%s's age is %d\n", a2.getName(), a2.getAge())
	fmt.Printf("Human %s can go %.2f km\n", a2.getName(), a2.getMaxKm())
}
