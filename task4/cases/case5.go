package cases

import (
	"fmt"
	"strings"
)

/*
InitCase5 ...
В ресторане действуют следующие правила:
	по понедельникам должна применяться скидка 10% на всё меню, потому что понедельник — день тяжёлый!
	по пятницам, если сумма чека превышает 10 000 руб., включается дополнительная скидка в размере 5%;
	если число гостей в одной компании превышает 5 человек, автоматически включается надбавка на обслуживание 10%.
Напишите программу, которая запрашивает день недели, число гостей и сумму чека и рассчитывает сумму к оплате.
*/
func InitCase5() {

	var (
		weekDay         string
		discount        int = 10
		service         int = 10
		checkSumm       int
		checkSummResult int
		countCustomers  int
	)

	fmt.Println("Введите сумму чека")
	fmt.Scan(&checkSummResult)
	fmt.Println("Введите число гостей")
	fmt.Scan(&countCustomers)
	fmt.Println("Введите день недели")
	fmt.Scan(&weekDay)

	switch strings.ToLower(weekDay) {
	case "понедельник", "пн":
		checkSumm = -checkSummResult / 100 * discount
	case "пятница", "пт":
		if checkSummResult > 10000 {
			checkSumm = checkSummResult / 100 * 5
		}
	}

	if countCustomers > 5 {
		checkSumm += checkSummResult / 100 * service
	}
	checkSummResult += checkSumm
	fmt.Println("Сумма к оплате:", checkSummResult)

}
