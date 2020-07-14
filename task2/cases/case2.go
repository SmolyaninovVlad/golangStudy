package cases

import (
	"fmt"
)

var (
	cost     int
	delivery int
	discount int
	total    int
)

func InitCase2() {
	cost = 6400
	delivery = 350
	discount = 700
	total = cost + delivery - discount

	fmt.Println("Стоимость товара: ", cost)
	fmt.Println("Стоимость доставки: ", delivery)
	fmt.Println("Размер скидки: ", discount)
	fmt.Println("Итого: ", total)
}
