package piscine

func AlphaCount(str string) int {

	counter := 0 // Создается переменная-счётчик для подсчета символов
	// Мог бы и сойти стандартный i-ый из цикла for но помоему область видимость не позволит ему
	// вернуть нужное значение

	for _, sign := range str {
		if sign >= 'a' && sign <= 'z' || sign >= 'A' && sign <= 'Z' {

			counter++
		}
	}

	// Второй способ
	// for i:= 0; i < str ; i++ {
	// 	if sign >= 'a' && sign <= 'z' || sign >= 'A' && sign <= 'Z'
	// 	counter++
	// }

	return counter
}
