package piscine

import "github.com/01-edu/z01"

func PrintNbrInOrder(n int) {
	if n == 0 {

		z01.PrintRune('0')
	}
	array := []int{}

	for n != 0 { // цикл while(n != 0)
		position := n % 10
		array = append(array, position) // append - сохранить в массив

		n = n / 10 // остекаем справа налево наш digit
	}

	len := 0
	for range array { // пройтись по всей длине массива
		len++
	}

	// BubbleSort
	for i := 1; i < len; i++ {
		for j := 1; j < len; j++ {
			if array[i-1] > array[i] {
				array[i-1], array[i] = array[i], array[i-1] // оказывается и так можно вместо func Swap()
			}
		}
	}

	for _, nb := range array { // пройтись по всему массиву
		z01.PrintRune(rune(nb + 48)) // принтит(оборавчивает в руну - каст)
	}

}
