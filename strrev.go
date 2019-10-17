package piscine

func StrRev(s string) string {
	temp := []rune(s)
	temp2 := []rune(s)
	len := 0
	for range s {
		len++
	}
	for i := 0; i < len; i++ {
		temp2[len-i-1] = temp[i]
	}
	return string(temp2)
}
