package piscine

func Sqrt(nb int) int {

	for i := 1; i <= nb; i++ {
		if i*i == nb {
			return i
		}
	}
	return 0 // po vsei vidimosti dlya sbrosa
}
