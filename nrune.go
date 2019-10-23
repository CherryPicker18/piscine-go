package piscine

func NRune(s string, n int) rune {
	arrOfRune := []rune(s)
	for i, value := range arrOfRune {
		if i == n-1 {
			return value
		}
	}
	return 0
}
