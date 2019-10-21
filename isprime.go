package piscine

func IsPrime(nb int) bool {

	// если число = 1  - false
	// если число = 0  - false
	// если число = -1 - false

	if nb == 1 || nb <= 0 {
		return false
	}

	isPrime := true

	for i := 2; i*i <= nb; i++ {
		if nb%i == 0 { // если делится на нацело - false
			isPrime = false
		}

	}

	return isPrime
}
