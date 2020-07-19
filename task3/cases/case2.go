package cases

import (
	"fmt"
)

var (
	ticketCost        int = 20
	total             int = 0
	currentPassangers int = 0
)

// InitCase2 ...
func InitCase2() {

	fmt.Println("Здравствуйте, вас приветствует симулятор маршрутки.")

	currentPassangers, total = station("Улица Программистов", currentPassangers, total)
	currentPassangers, total = station("Проспект Алгоритмов", currentPassangers, total)
	currentPassangers, total = station("Ул. им. Функции", currentPassangers, total)
	currentPassangers, total = station("Полиморфная улица", currentPassangers, total)

	printResult(float64(total))
}

func printResult(total float64) {
	var (
		driverSalary  float64 = total * 0.25
		gasOutcome    float64 = total * 0.2
		taxOutcome    float64 = total * 0.2
		repairOutcome float64 = total * 0.2
		income        float64 = total - driverSalary - gasOutcome - taxOutcome - repairOutcome
	)
	fmt.Println("Всего заработали:", total)
	fmt.Println("Зарплата водителя:", driverSalary)
	fmt.Println("Расходы на топливо:", gasOutcome)
	fmt.Println("Налоги:", taxOutcome)
	fmt.Println("Расходы на ремонт машины:", repairOutcome)
	fmt.Println("Итого доход:", income)
}
func station(stationName string, currentPassangers int, total int) (int, int) {
	fmt.Println("-----------Едем---------")
	arriving("Улица Программистов", currentPassangers)
	currentPassangers = passangersLeft(currentPassangers)
	currentPassangers, total = passangersEntered(currentPassangers, total)
	departing("Улица Программистов", currentPassangers)
	return currentPassangers, total
}
func passangersLeft(count int) int {
	var passangersLeft int
	fmt.Println("Сколько пассажиров вышло на остановке?")
	fmt.Scan(&passangersLeft)
	count -= passangersLeft
	return count
}
func passangersEntered(count int, total int) (int, int) {
	var passangersEntered int
	fmt.Println("Сколько пассажиров зашло на остановке?")
	fmt.Scan(&passangersEntered)
	count += passangersEntered
	total += passangersEntered * 20
	return count, total
}
func arriving(station string, passangers int) {
	fmt.Println("Прибываем на остановку ", station, ". В салоне пассажиров:", passangers)
}
func departing(station string, passangers int) {
	fmt.Println("Отправляемся с остановки ", station, ". В салоне пассажиров:", passangers)
}
