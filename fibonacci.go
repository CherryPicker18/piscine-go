package piscine

func Fibonacci(index int) int {

	if index < 0 { // nothing more
		return -1
	}
	if index == 0 { // 0 + 0 = 0
		return 0
	}
	if index == 1 { // 0 + 1 = 1
		return 1
	}

	return Fibonacci(index-1) + Fibonacci(index-2)
}
