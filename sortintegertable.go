package piscine

func swapit(a *int, b *int) {
	c := *b
	*b = *a
	*a = c
}

func SortIntegerTable(table []int) {
	temp := table
	n := len(temp)
	for i := 0; i < n; i++ {
		for j := n - 1; j > i; j-- {
			if temp[j] < temp[j-1] {
				swapit(&temp[j], &temp[j-1])
			}
		}
	}
}
