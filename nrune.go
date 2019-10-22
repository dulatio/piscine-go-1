package piscine

func RuneArrayLength(array []rune) int {
	length := 0
	for range table {
		length++
	}
	return length
}

func NRune(s string, n int) rune {
	str := []rune(s)
	if n > RuneArrayLength(str)-1 {
		return 0
	}
	return str[n-1]
}
