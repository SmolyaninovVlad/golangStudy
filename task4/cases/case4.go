package cases

import (
	"fmt"
)

/*
InitCase4 ...
Напишите программу, которая запрашивает у пользователя три числа и выводит количество чисел, которые больше, либо равны 5.
*/
func InitCase4() {
	var (
		value1 int
		value2 int
		value3 int
		count  int
	)

	fmt.Println("Введите три числа и программа сообщит вам сколько чисел большее чем 5")

	fmt.Println("введите первое число")
	fmt.Scan(&value1)
	fmt.Println("введите второе число")
	fmt.Scan(&value2)
	fmt.Println("введите третье число")
	fmt.Scan(&value3)

	if value1 > 5 {
		count++
	}

	if value2 > 5 {
		count++
	}
	if value3 > 5 {
		count++
	}

	fmt.Println("Чисел больше 5:", count)
}
