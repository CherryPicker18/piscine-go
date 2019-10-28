package main

import (
	"os"

	"github.com/01-edu/z01"
)

func printStr(str string) {
	arrayStr := []rune(str)

	var count int

	for index := range arrayStr {
		count = index
	}

	for i := 0; i <= count; i++ {
		z01.PrintRune(arrayStr[i])
	}

	z01.PrintRune('\n')
}

func isEven(nbr int) bool {
	if nbr == 1 {
		return true
	} else {
		return false
	}
}

func main() {

	a := os.Args[1:]
	var k int

	for index := range a {
		k = index
	}

	lengthOfArg := k % 2

	EvenMsg := "I have an odd number of arguments"
	OddMsg := "I have an even number of arguments"

	if isEven(lengthOfArg) {
		printStr(OddMsg)
	} else {
		printStr(EvenMsg)
	}
}
