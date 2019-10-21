package piscine

func IterativePower(nb int, power int) int {
	answer := nb
	if power < 0 {
		return 0
	} else if power == 0 {
		return 1
	}
	for i := 2; i <= power; i++ {
		answer *= nb
	}
	return answer
}
