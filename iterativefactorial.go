package piscine 

import "github.com/01-edu/z01"

func IterativeFactorial(nb int)int {

	if nb >= 0 && nb <= 20 {

	Fact:= 1
	i:= 1
	for ; i <= nb; Fact *= i  {

		
		i++
	} else {
		return 0
	} 

	}
		
		
}