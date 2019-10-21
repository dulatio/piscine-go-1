package piscine

func FindNextPrime(nb int) int {
	var i int
	if nb <= 1 {
		return 2
	}
	for i = nb; i <= 2*nb; i++ {
		if IsPrime(i) == true {
			break
		}
	}
	return i
}
