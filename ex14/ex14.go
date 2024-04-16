package main

// Разработать программу, которая в рантайме способна определить тип переменной: int, string, bool, channel
// из переменной типа interface{}.

import (
	"fmt"
	"reflect"
)

func main() {

	var t1 interface{} = make(chan string) // создаем переменную с типом interface и записываем в нее любое значение

	fmt.Printf("t1: %s\n", reflect.TypeOf(t1)) // получаем тип значения с помощью reflect.TypeOf

	switch t1.(type) { // получаем тип значения через switch-case.
	//В этом случае для channel нужно задавать все возможные варианты, т.к. тип channel interface не будет принимать по умолчанию  channel с любыми типами

	case int:
		fmt.Println("it`s int")
	case string:
		fmt.Println("it`s string")
	case bool:
		fmt.Println("it`s bool")
	case chan int, chan string, chan bool, chan interface{}:
		fmt.Println("it`s channel")
	default:
		fmt.Println("unknown type")
	}
}
