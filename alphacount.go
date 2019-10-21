package piscine

func AlphaCount(str string) int {
	answer := 0
	strc := []rune(str)
	for _, char := range strc {
		if (char >= 'a' && char <= 'z') || (char >= 'A' && char <= 'Z') {
			answer++
		}
	}
	return answer
}
