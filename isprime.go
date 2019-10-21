package piscine

func IsPrime(nb int) bool {

	// если число = 1  - false
	// если число = 0  - false
	// если число = -1 - false

	// нужен лог флажок ?
	// isPrime := false

	if (nb%1 == 0) && (nb%nb == 0) {

		for i = 0; i < nb; i++ {
			nb%i == 0
			return true
		}

		if nb == 1 {
			return false
		}
		if nb <= 0 {
			return false
		}
	}

}
