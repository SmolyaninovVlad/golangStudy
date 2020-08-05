package cases

import (
	"fmt"
)

var value int

/*
InitCase2 ...
Напишите программу, которая запрашивает у пользователя три числа и сообщает, есть ли среди них число, большее, чем 5.
*/
func InitCase2() {
	var values []int

	fmt.Println("Введите три числа и программа сообщит вам есть ли среди них число большее чем 5")

	bigValue := 0

	for i := 0; i < 3; i++ {
		fmt.Println("введите число")
		fmt.Scan(&value)
		values = append(values, value)
	}

	for _, value := range values {
		if value > 5 {
			bigValue++
		}
	}

	if bigValue > 0 {
		fmt.Println("Есть число больше 5")
	} else {
		fmt.Println("Нет чисел значения которых больше 5")
	}

}
