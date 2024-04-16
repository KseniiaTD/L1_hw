package main

// Имеется последовательность строк - (cat, cat, dog, cat, tree) создать для нее собственное множество.

import "fmt"

func main() {
	arr := []string{"cat", "cat", "dog", "cat", "tree"} //задаем слайс строк
	m := make(map[string]struct{})                      // создаем пустое множество
	for _, i := range arr {                             // обходим слайс, записывая значения в множество
		m[i] = struct{}{}
	}

	fmt.Println(m)
}
