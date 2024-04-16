package main

// Реализовать пересечение двух неупорядоченных множеств.

import "fmt"

func intersection1(set1, set2 map[string]struct{}) map[string]struct{} {
	//функция для пересечения множеств, где считаем кол-во повторений для каждого из значений обоих множеств
	var set3 = make(map[string]int)        // задаем множество 3
	var setRes = make(map[string]struct{}) // задаем результирующее множество

	for key := range set1 { //подсчитываем количество значений из первого множества в третьем множестве
		set3[key]++
	}
	for key := range set2 { //подсчитываем количество значений из второго множества в третьем множестве
		set3[key]++
	}

	for key, cnt := range set3 { // записываем в результирующее множества только те значеня из множества 3, у которых кол-во повторений = 2
		// это значит, что значение встретилось в 1 и 2 множестве
		if cnt == 2 {
			setRes[key] = struct{}{}
		}
	}

	return setRes //возвращаем результирующее множество
}

func intersection2(set1, set2 map[string]struct{}) map[string]struct{} {
	//функция для пересечения множеств, где проверяем наличие значения из 1 множества во 2 множестве
	var setRes = make(map[string]struct{}) // создаем результирующее множество
	for key := range set1 {                // в цикле обходим множество 1
		_, exist := set2[key] // проверяем наличие значения во втором множестве
		if exist {            // если значение найдено, добавляем в результирующее множество
			setRes[key] = struct{}{}
		}
	}

	return setRes //возвращаем результирующее множество
}

func main() {
	var set1 = map[string]struct{}{ //задаем множество 1
		"John":    {},
		"Sam":     {},
		"Lilo":    {},
		"Barbara": {},
	}
	var set2 = map[string]struct{}{ // задаем множество 2
		"David":   {},
		"John":    {},
		"Fred":    {},
		"Kate":    {},
		"Samanta": {},
		"Barbara": {},
	}

	fmt.Println(intersection1(set1, set2)) // вариант 1 пересечения множеств
	fmt.Println(intersection2(set1, set2)) // вариант 2 пересечения множеств
}
