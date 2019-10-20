package piscine

func IterativeFactorial(nb int) int {
	fuck := 1
	if nb >= 0 && nb <= 20 {
		for i := 1; i <= nb; i++ {
			fuck = fuck * i
		}
	} else {
		return 0
	}
}
