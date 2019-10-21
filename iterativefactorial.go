package piscine

func IterativeFactorial(nb int) int {
	answer := 1
	for i := 2; i <= nb; i++ {
		answer *= nb
	}
	return answer
}
