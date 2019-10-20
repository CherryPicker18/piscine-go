package piscine

func IterativeFactorial(nb int) int {
	if nb >= 0 && nb <= 20 {
		Fact := 1
		i := 1
		for ; i <= nb; i++ {
			Fact *= i
		}
		return (Fact)
	} else {
		return 0
	}

}
