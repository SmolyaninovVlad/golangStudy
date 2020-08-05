package cases

import (
	"fmt"
)

var (
	score    int
	minScore int = 275
)

/*
InitCase1 ...
Сумма проходных баллов в институт составляет 275 баллов.
Напишите программу, которая запрашивает у пользователя результаты ЕГЭ по трём экзаменам и проверяет, поступил ли он в институт или нет.
*/
func InitCase1() {

	fmt.Println("Введите сумму ваших баллов по ЕГЭ по 3-м экзаменам")

	fmt.Scan(&score)
	if score < minScore {
		fmt.Println("Колличество баллов не достаточное для поступления в институт")
	} else {
		fmt.Println("Вы поступили в институт")
	}
}
