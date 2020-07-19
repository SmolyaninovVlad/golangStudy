package cases

import (
	"fmt"
)

var (
	currentHeight int = 100 //1 метр -нач. высота
	atePerNight   int = 20
	growthPerDay  int = 50
	daysCount     int
	goalHeight    int = 300 //цель
	firstPartDays int
)

// InitCase4 Злостные вредители
func InitCase4() {
	fmt.Println("Здравствуйте, вас приветствует программа расчёта высоты бамбука")

	fmt.Println("Введите начальную высоту бамбука (см.):")
	fmt.Scan(&currentHeight)
	fmt.Println("Введите скорость роста бамбука в день (см.):")
	fmt.Scan(&growthPerDay)
	fmt.Println("Введите скорость поедания бамбука за ночь (см.):")
	fmt.Scan(&atePerNight)
	fmt.Println("Введите количество дней для расчёта высоты бамбука (1 часть задания):")
	fmt.Scan(&firstPartDays)
	fmt.Println("Введите целевую высоту взрослого бамбука (2 часть задания):")
	fmt.Scan(&goalHeight)

	oneDayGrowth := (growthPerDay - atePerNight)

	// firstPart
	height := currentHeight + oneDayGrowth*(firstPartDays-1) + growthPerDay/2
	fmt.Println("В середине", firstPartDays, "дня высота бамбука будет состовлять", height, "см")

	// secondPart
	daysCount = (goalHeight-currentHeight-growthPerDay)/oneDayGrowth + 1
	fmt.Println("Для достижения высоты", goalHeight, "см. бамбуку понадобиться рости", daysCount, "дней")
}
