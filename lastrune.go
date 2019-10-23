package piscine

func LastRune(s string) rune {

	arror := []rune(s) // Создаётся слайс рун из строки

	len := 0 // иниц перем для подсчета длины

	for range arror {

		len++
	}

	for index, value := range arror { // Пробежка по значениям массива
		if index == len-1 {
			return value
		}
	}

	return 42 // исчерпывающий ответ
}
