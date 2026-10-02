package main

import (
	"errors"
	"strings"
	"time"
	"unicode/utf8"
)

// объявление текстов ошибок
var (
	ErrByLen = errors.New("исправь свой ответ, а лучше ложись поспать")
)

// функция, выводящая текущий день недели
func currentDayOfTheWeek() string {
	data := time.Now()
	dayOfWeek := data.Weekday()
	switch dayOfWeek{
	case time.Saturday:
		return "Суббота"
	case time.Monday:
		return "Понедельник"
	case time.Tuesday:
		return "Вторник"
	case time.Wednesday:
		return "Среда"
	case time.Thursday:
		return "Четверг"
	case time.Friday:
		return "Пятница"
	case time.Sunday:
		return "Воскресенье"
	default:
		return "Ошибка"
	}
}

// выводит день или ночь по времени суток
func dayOrNight() string {
	data := time.Now()
	timeDay := data.Hour()
	if timeDay >= 10 && timeDay <= 22 {
		return "День"
	}
	return "Ночь"
}

// выводит сколько остлось дней до пятницы
func nextFriday() int {
	dayOfWeek := currentDayOfTheWeek()
	switch dayOfWeek {
	case "Воскресенье":
		return 5
	case "Понедельник":
		return 4
	case "Вторник":
		return 3
	case "Среда":
		return 2
	case "Четверг":
		return 1
	case "Пятница":
		return 0
	case "Суббота":
		return 6
	default:
		return -1
	}
}

// определяет прав пользователь или нет в своем предположении о том, какой сегодня день недели
func CheckCurrentDayOfTheWeek(answer string) bool {
	hz := currentDayOfTheWeek()
	if strings.ToLower(answer) == strings.ToLower(hz) {
		return true
	}
	return false
}

// то же самое, что и предыдущая функция, но день или ночь
func CheckNowDayOrNight(answer string) (bool, error) {
	hz := dayOrNight()
	if strings.ToLower(answer) == strings.ToLower(hz) {
		return true, nil
	} else if utf8.RuneCountInString(answer) < 4 || utf8.RuneCountInString(answer) > 4 {
		return false, ErrByLen
	}
	return false, nil
}
