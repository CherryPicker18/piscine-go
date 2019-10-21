package piscine

func IterativePower(nb int, power int) int {
	result := nb // in result we save our digit

	if power < 0 {
		return 0
	}
	if power == 0 {
		return 1
	}
	for p := 2; p <= power; p++ { // go in circle from 2 cuz upper 2 if's :d
		result = result * nb // 2
	}

	return result
}
