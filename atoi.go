package piscine

func intfor3(r rune) int {
	temp2 := -1
	for temp := '0'; temp <= r; temp++ {
		temp2++
	}
	return temp2
}

func Atoi(s string) int {
	bool1 := false
	arrayStr := []rune(s)
	n := 0
	temp := 0
	for range arrayStr {
		n++
	}
	if n != 0 && arrayStr[0] == '-' {
		bool1 = true
		temp++
	}
	ans := 0
	for i := temp; i < n; i++ {
		if arrayStr[i] < '0' || arrayStr[i] > '9' {
			return 0
		} else {
			ans *= 10
			ans += intfor(arrayStr[i])
		}
	}
	if bool1 == true {
		ans = ans * -1
	}
	return ans
}
