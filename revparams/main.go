package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	a := os.Args
	k := 0

	for index := range a {
		k = index
	}

	for i := k; i >= 1; i-- {
		for _, letter := range a[i] {
			z01.PrintRune(letter)
		}
		z01.PrintRune('\n')
	}
}
