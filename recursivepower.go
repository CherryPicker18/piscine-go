package piscine

func RecursivePower(nb int, power int) int {
	if power < 0 { // if power (-1) +/-; +/-
		return 0
	}
	if power == 0 { // anyway digit with power 0 = 1
		return 1
	}
	return nb * RecursivePower(nb, power-1) // RP(2,3)  2 * 4 = 8(2 * 2 * 2)
}
