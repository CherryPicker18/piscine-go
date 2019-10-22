package piscine

func TrimAtoi(s string) int {

	runes := []rune(s)    // Создан массив рун из строки
	isFoundDigit := false // Создан логический флажок
	result := 0           // иниц переменная для результата
	sign := 1             // ициц переменная для символа

	for _, r := range runes {
		if !isFoundDigit {

			if r == '-' {
				sign = -1
			} else if r == '+' {
				sign = 1
			}
		}

		if r >= '0' && r <= '9' {
			isFoundDigit = true
			result = 10*result + int(r-'0')
		}
	}

	return sign * result
}
