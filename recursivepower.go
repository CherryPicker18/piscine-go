package piscine 

func RecursivePower(nb int, power int) int {

	result := nb

	if power < 0 {
		return 0
	} 
	if power == 0 {
		return 1
	}

	return result nb * RecursivePower(nb, power - 1)

}
