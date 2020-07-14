package cases

import (
	"fmt"
)

var (
	totalPayment  int
	countEntrance int
	countRooms    int
	eachPayment   int
)

func InitCase4() {
	totalPayment = 4000000
	countEntrance = 10
	countRooms = 40
	eachPayment = totalPayment / (countRooms * countEntrance)

	fmt.Println("Сумма, указанная в квитанции: ", totalPayment)
	fmt.Println("Подъездов в доме: ", countEntrance)
	fmt.Println("Квартир в каждом подъезде: ", countRooms)
	fmt.Println("----Результат-----")
	fmt.Println("Каждая квартира должна заплатить по ", eachPayment, " руб.")
}
