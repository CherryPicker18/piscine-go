package piscine

func IterativePower(nb int, power int) int {

	result := nb // there we save digit
	// 2 * 2 = 4 * 2 (digit = 2 ; if power = 3)
	if power < 0 {
		return 0
	}

	if power == 0 {
		return 1
	}

	for i := 2; i <= power; i++ { // begig from 2 cuz we don't need 0 and 1
		digit = nb * digit // upper 2 if
	}

	return digit
}
