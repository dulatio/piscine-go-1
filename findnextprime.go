package piscine

func FindNextPrime(nb int) int {
	var i int
	if nb == 0 {
		return 1
	}
	for i = nb; i <= 2*nb; i++ {
		if IsPrime(i) == true {
			break
		}
	}
	return i
}
