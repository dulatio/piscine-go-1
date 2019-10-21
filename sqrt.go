package piscine

func Sqrt(nb int) int {
	answer := 1
	for i := 1; i*i <= nb; i++ {
		answer = i
	}
	if answer*answer == nb {
		return answer
	} else {
		return 0
	}
}
