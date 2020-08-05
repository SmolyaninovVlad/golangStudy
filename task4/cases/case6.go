package cases

import (
	"fmt"
)

// var (
// 	students        []int
// 	studentsCurrent []int
// 	groups          [][]int
// )

/*
InitCase6 ...
Перед старостой группы стоит задача разделить весь курс, состоящий из N студентов, на K групп.
Напишите программу, которая поможет старосте сделать это: он вводит N, K и порядковый номер студента, а программа определяет, в какую группу он попадает.
Подсказка: в одну группу могут попадать студенты из разных частей списка.
*/
func InitCase6() {
	var (
		studentsCount int
		groupsCount   int
		studentNumber int
	)
	fmt.Println("Колличество студентов")
	fmt.Scan(&studentsCount)
	fmt.Println("Колличество групп")
	fmt.Scan(&groupsCount)
	fmt.Println("Введите порядковый номер студента")
	fmt.Scan(&studentNumber)

	// Узнаем сколько студентов в 1 группе
	studentsInGroup := studentsCount / groupsCount
	if studentsCount%groupsCount != 0 {
		studentsInGroup++
	}

	// Порядковый номер студента меньше чем человек в любой группе то он помещается в 1 группу
	if studentNumber <= studentsInGroup {
		fmt.Println("Выбранный студент распределяется в 1 группу")
		return
	}

	// Считаем в какую по счёту группу попадает студент
	studentGroup := studentNumber / studentsInGroup
	if studentNumber%studentsInGroup != 0 {
		studentGroup++
	}

	fmt.Println("Выбранный студент распределяется в", studentGroup, "группу")

	// createStudents(studentsCount)
	// createGroups(groupsCount)
	// fillGroups()

	// fmt.Println("Все студенты распределены на ", groupsCount, " групп")
	// fmt.Println("Сформированные группы", groups)
}

// Заполнение массива студентов их порядковыми номерами
// func createStudents(studentsCount int) {
// 	for i := 1; i <= studentsCount; i++ {
// 		students = append(students, i)
// 	}
// }

// func createGroups(groupsCount int) {
// 	for i := 1; i <= groupsCount; i++ {
// 		groups = append(groups, []int{})
// 	}
// }

// // Функция распределения студентов по группам
// func fillGroups() {
// 	studentsCurrent = make([]int, len(students))
// 	copy(studentsCurrent, students)
// 	for range students {
// 		index := requestStudentNumber()
// 		//Определение группы с наименьшим колличеством студентов
// 		groupIndex := getGroup()
// 		//Добавление студента в группу
// 		groups[groupIndex] = append(groups[groupIndex], studentsCurrent[index])
// 		//Удаление перемещённого студента из списка нераспределённых
// 		studentsCurrent = removeIndex(index)
// 	}
// }

// // Получение индекса по введённому значению
// func requestStudentNumber() int {
// 	var studentNumber int
// 	fmt.Println("Порядковые номера нераспределённых студентов:", studentsCurrent)
// 	fmt.Println("Введите порядковый номер студента")
// 	fmt.Scan(&studentNumber)
// 	index, err := findIndex(studentNumber)
// 	if err != nil {
// 		fmt.Println(err)
// 		index = requestStudentNumber()
// 		return index
// 	}
// 	return index
// }

// // Получение минимальной по численности студентов группы
// func getGroup() int {
// 	minLenIndex := 0
// 	minLenGroup := groups[minLenIndex]

// 	if len(groups[minLenIndex]) == 0 {
// 		return minLenIndex
// 	}

// 	for index, group := range groups {
// 		if len(minLenGroup) > len(group) {
// 			minLenGroup = group
// 			minLenIndex = index
// 		}
// 	}
// 	return minLenIndex
// }
// func findIndex(needValue int) (int, error) {
// 	var err = errors.New("Нет студента с таким порядковым номером")
// 	for index, value := range studentsCurrent {
// 		if value == needValue {
// 			return index, nil
// 		}
// 	}
// 	return -1, err
// }
// func removeIndex(index int) []int {
// 	return append(studentsCurrent[:index], studentsCurrent[index+1:]...)
// }
