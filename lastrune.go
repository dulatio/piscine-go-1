package piscine

func LastRune(s string) rune {
	str := []rune(s)
	return str[RuneArrayLength(str)-1]
}
