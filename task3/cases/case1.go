package cases

import (
	"fmt"
)

var (
	planet string
	star   string
	ranger string
	money  int
)

// InitCase1 ...
func InitCase1() {

	fmt.Println("Здравствуйте, это небольшой текстовый квест.")

	fmt.Println("Введите название планеты:")
	fmt.Scan(&planet)
	fmt.Println("Введите название звёздной системы:")
	fmt.Scan(&star)
	fmt.Println("Введите имя рэйнджера:")
	fmt.Scan(&ranger)
	fmt.Println("Введите количество вознаграждения:")
	fmt.Scan(&money)

	fmt.Println("\n\n Как вам,", ranger, ", известно, мы - раса мирная, поэтому на наших военных кораблях используются наемники с других\n",
		"планет. Система набора отработана давным-давно. Обычно мы приглашаем на наши корабли людей с планеты", planet, " системы ", star)

	fmt.Println("\n\n Но случилась неприятность - в связи с большими потерями, в последнее время, престиж профессии сильно упал, и\n",
		"теперь не так-то просто завербовать призывников. Главный комиссар планеты", planet, " впрочем, отлично справлялся,\n",
		"но недавно его наградили орденом Сутулого с закруткой на спине, и его на радостях парализовало! Призыв под\n",
		"угрозой срыва!")

	fmt.Println("\n\n", ranger, "вы должны прибыть на планету", planet, "системы", star, "и помочь выполнить план призыва. Мы готовы\n",
		"выплатить вам премию в", money, "кредитов за эту маленькую услугу.")

}
