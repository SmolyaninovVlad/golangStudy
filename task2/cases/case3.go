package cases

import (
	"fmt"
)

var (
	workTime     float64
	orderTime    float64
	serviceTime  float64
	countClients float64
)

func InitCase3() {
	fmt.Println("Эта программа рассчитает, сколько клиентов успеет обслужить кассир за смену.")
	// имиация ввода данных с клавиатуры
	workTime = 480
	fmt.Println("Введите длительность смены в минутах: ", workTime)
	orderTime = 2
	fmt.Println("Сколько минут клиент делает заказ? ", orderTime)
	serviceTime = 4
	fmt.Println("Сколько минут кассир собирает заказ? ", serviceTime)

	// расчёт кол-ва клиентов
	fmt.Println("-----Считаем-----")
	countClients = workTime / (orderTime + serviceTime)
	fmt.Println("За смену длиной ", workTime, " минут кассир успеет обслужить ", countClients, " клиентов.")
}
