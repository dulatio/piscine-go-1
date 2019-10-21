package piscine

func IterativeFactorial(nb int) int {
	answer := 1
	if nb < 0 {
		return 0
	}
	for i := 2; i <= nb; i++ {
		answer *= i
	}
	return answer
}
