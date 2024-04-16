package main

// Реализовать паттерн «адаптер» на любом примере.

//Тут будем считать бонусы сотрудникам по формуле : факт продаж * личный процент / 100
//Но  кроме сотрудников бонус будем считать и директору, у которого нет своего метода факта продаж
// для директора создадим адаптер, рассчитывающий факт продаж как среднее значение от факта продаж закрепленных сотрудников

import "fmt"

type BonusCalc interface { // интерфейс для рассчета бонуса
	CalculateFact() float32 // получаем факт продаж
	GetPercentEmp() int     // получаем личный процент
}

type Employee struct { // структура Сотрудник
	name    string  //имя сотрудника
	fact    float32 // факт продаж
	percent int     // личный процент
}

func (e *Employee) CalculateFact() float32 { // возвращаем факт продаж
	return e.fact
}

func (e *Employee) GetPercentEmp() int { // возвращаем личный процент
	return e.percent
}

type Director struct { // структура Директор
	name      string     // имя директора
	employees []Employee // слайс сотрудников, закрепленных за директором
	percent   int        // личный процент
}

func (d *Director) CalculateEmp() int { // количество сотрудников, закрепленных за директором
	return len(d.employees)
}

func (d *Director) CalculateSumFact() float32 { // Количество продаж отделения
	var sum float32
	for _, i := range d.employees { // суммируем факты продаж всех сотрудников отделения
		sum += i.fact
	}
	return sum
}

func (d *Director) GetPercentDir() int { // получаем личный процент для директора
	return d.percent
}

type DirAdapter struct { // структура Адаптер для директора
	director *Director
}

func (d *DirAdapter) CalculateFact() float32 { // получаем факт продаж для директора как для сотрудника
	cnt := d.director.CalculateEmp() // делим общее кол-во продаж отделения на кол-во сотрудников в отделении
	sum := d.director.CalculateSumFact()
	return sum / float32(cnt)
}

func (d *DirAdapter) GetPercentEmp() int { // получаем факт продаж для директора как для сотрудника
	return d.director.GetPercentDir()
}

func printBonus(calc BonusCalc) { // рассчитываем бонус для сотрудника через интерфейс
	fmt.Println("Бонус:", calc.CalculateFact()*float32(calc.GetPercentEmp())/float32(100))
}

func main() {
	e1 := Employee{name: "Bob", fact: 1045.7, percent: 5} // задаем 3 сотрудников
	e2 := Employee{name: "Kate", fact: 3032.6, percent: 5}
	e3 := Employee{name: "Mark", fact: 4515.2, percent: 5}
	d1 := Director{name: "John", employees: []Employee{e1, e2, e3}, percent: 7} // задаем директора
	adapterD := DirAdapter{director: &d1}                                       // создаем адаптер для директора, для рассчета бонуса через интерфейс

	// рассчитываем бонус для сотрудников и директора, используя единый интерфейс
	printBonus(&e1)
	printBonus(&e2)
	printBonus(&e3)
	printBonus(&adapterD)
}
