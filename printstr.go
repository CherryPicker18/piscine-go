package piscine

import (
	"github.com/01-edu/z01"
)

func PrintStr(str string) {

	r := []rune(str)
	a := 0
	for index := range str {
		a = index
	}
	for i := 0; i <= a; i++ {
		z01.PrintRune(r[i])
	}
}
