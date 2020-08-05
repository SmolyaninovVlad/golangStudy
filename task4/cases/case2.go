package cases

import (
	"fmt"
)

/*
InitCase2 ...
Напишите программу, которая запрашивает у пользователя три числа и сообщает, есть ли среди них число, большее, чем 5.
*/
func InitCase2() {
	var (
		value1 int
		value2 int
		value3 int
	)

	fmt.Println("Введите три числа и программа сообщит вам есть ли среди них число большее чем 5")

	fmt.Println("введите первое число")
	fmt.Scan(&value1)
	fmt.Println("введите второе число")
	fmt.Scan(&value2)
	fmt.Println("введите третье число")
	fmt.Scan(&value3)

	if value1 > 5 || value2 > 5 || value3 > 5 {
		fmt.Println("Есть число больше 5")
	} else {
		fmt.Println("Нет чисел больше  5")
	}

}
