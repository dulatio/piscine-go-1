package piscine

func IsNumeric(s string) bool {
	str := []rune(s)
	for i := 0; i < RuneArrayLength(str); i++ {
		if str[i] > '9' || str[i] < '0' {
			return false
		}
	}
	return true
}
