package piscine 

func IterativeFactorial(nb int)int {

	if nb >= 0 && nb <= 20 {

	Fact := 1

	for i:= 1 ; i <= nb ; i++ {
		Fact = Fact * i 
	}
		
		
	} else {
		return 0
	}		
}

