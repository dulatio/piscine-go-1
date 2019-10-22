package piscine

func RuneArrayLength(array []rune) int {
	length := 0
	for range array {
		length++
	}
	return length
}

func NRune(s string, n int) rune {
	str := []rune(s)
	if n > RuneArrayLength(str) || n < 1 {
		return 0
	}
	return str[n-1]
}
