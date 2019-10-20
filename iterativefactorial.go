package piscine

func IterativeFactorial(nb int) int {

	Fact := 1

	if nb >= 0 && nb <= 20 {

		for i := 1; i <= nb; i++ {
			Fact = Fact * i
		}

	} else {
		return 0
	}
}
