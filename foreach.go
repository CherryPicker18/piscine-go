package piscine

func ForEach(f func(int), arr []int) {

	for _, value := range arr {
		f(value)
	}
}
