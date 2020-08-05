package cases

import (
	"fmt"
)

/*
InitCase4 ...
Напишите программу, которая запрашивает у пользователя три числа и выводит количество чисел, которые больше, либо равны 5.
*/
func InitCase4() {
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
		fmt.Println("Чисел значения которых больше 5:", bigValue)
	} else {
		fmt.Println("Нет чисел значения которых больше 5")
	}

}
