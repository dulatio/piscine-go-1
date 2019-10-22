package piscine

func NRune(s string, n int) rune {
	str := []rune(s)
	return str[n-1]
}
