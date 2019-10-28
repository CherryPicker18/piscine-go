package piscine

import "github.com/01-edu/z01"

func PrintStr(str string) {
	arrOfRunes := []rune(str)

	length := 0

	for range str {
		length++
	}

	for _, value := range arrOfRunes {
		z01.PrintRune(value)
	}

	z01.PrintRune('\n')
}

func main() {

	PrintStr("Hello")

}
