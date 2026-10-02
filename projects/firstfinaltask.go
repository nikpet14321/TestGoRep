package main

import (
	"fmt"
)

func main() {
	// объявление переменных
	var (
		a string = "1. -"
		flag1 bool = true
		b string = "2. -"
		flag2 bool = true
		c string = "3. -"
		flag3 bool = true
		d string = "4. -"
		flag4 bool = true
		e string = "5. -"
		flag5 bool = true
		sum int = 5
	)
	// бесконечный цикл, который заканчивается посредством ввода слова "конец" с помощь break
	for {
		// вводимые переменные
		var(
			word string
			num int
		)
		n, err := fmt.Scanln(&word, &num) // здесь есть err для того, чтобы считать сколько вводимых значений ввел пользователь
		if word == "" && err != nil {
			continue
		}
		// типо функции. Выводит текущую очередь
		if word == "очередь" {
			fmt.Println(a)
			fmt.Println(b)
			fmt.Println(c)
			fmt.Println(d)
			fmt.Println(e)
			continue
		}
		// Выводит текущее количество свободных и занятых мест
		if word == "количество" {
			s1 := fmt.Sprintf("Осталось свободных мест: %d", sum)
			s2 := fmt.Sprintf("Всего человек в очереди: %d", 5 - sum)
			fmt.Println(s1)
			fmt.Println(s2)
			continue
		}
		// заканчивает цикл
		if word == "конец" {
			fmt.Println(a)
			fmt.Println(b)
			fmt.Println(c)
			fmt.Println(d)
			fmt.Println(e)
			break
		}
		// сама логика записи на место и проверки занятости места для каждого номера очереди(надо было использовать switch, но что-то я решил написать через if else)
		if word != "" &&  err == nil && n == 2 {
			if num >= 1 && num <= 5 {
				if flag1 == false && flag2 == false && flag3 == false && flag4 == false && flag5 == false {
					s4 := fmt.Sprintf("Запись на место номер %d невозможна: очередь переполнена", num)
					fmt.Println(s4)
				} else {
					if num == 1 {
					if flag1 {
						a = fmt.Sprintf("1. %s", word)
						flag1 = false
						sum -= 1
						continue
					} else {
						fmt.Println("Запись на место номер 1 невозможна: место уже занято")
						continue
					}
				}
				if num == 2 {
					if flag2 {
						b = fmt.Sprintf("2. %s", word)
						flag2 = false
						sum -= 1
						continue
					} else {
						fmt.Println("Запись на место номер 2 невозможна: место уже занято")
						continue
					}
				}
				if num == 3 {
					if flag3 {
						c = fmt.Sprintf("3. %s", word)
						flag3 = false
						sum -= 1
						continue
					} else {
						fmt.Println("Запись на место номер 3 невозможна: место уже занято")
						continue
					}
				}
				if num == 4 {
					if flag4 {
						d = fmt.Sprintf("4. %s", word)
						flag4 = false
						sum -= 1
						continue
					} else {
						fmt.Println("Запись на место номер 4 невозможна: место уже занято")
						continue
					}
				}
				if num == 5 {
					if flag5 {
						e = fmt.Sprintf("5. %s", word)
						flag5 = false
						sum -= 1
						continue
					} else {
						fmt.Println("Запись на место номер 5 невозможна: место уже занято")
						continue
					}
				}
				}
				
			} else {
				s3 := fmt.Sprintf("Запись на место номер %d невозможна: некорректный ввод", num)
				fmt.Println(s3)
				continue
			}
		} else {
			fmt.Println("некорректный ввод")
			continue
		}
	}
}