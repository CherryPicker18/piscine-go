package piscine

import "github.com/01-edu/z01"

func checkit(v int) {
	check := '0'
	if v == 0 {
		z01.PrintRune(check)
		return
	}
	for i := 1; i <= v%10; i++ {
		check++
	}
	for i := -1; i >= v%10; i-- {
		check++
	}
	if v/10 != 0 {
		checkit(v / 10)
	}
	z01.PrintRune(check)
	return
}

func PrintNbr(n int) {
	if n < 0 {
		z01.PrintRune('-')
	}
	checkit(n)
}
