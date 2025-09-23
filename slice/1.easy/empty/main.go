package main

import "fmt"

// Не запуская, подумать, что будет выведено на экран (true/false).
// Проверить и разобраться, почему так.
// Написать ответ подробно.

func main() {
	var list []int
	fmt.Println(list == nil)

	list = []int{}
	fmt.Println(list == nil)
}
