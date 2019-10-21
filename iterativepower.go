package piscine

func IterativePower(nb int, power int) int {

	if power < 0 {
		return 0
	}

	if power == 0 {
		return 1
	}

	digit := nb

	for i := 0; i < power; i++ {
		digit = nb * digit
	}

	return digit
}
