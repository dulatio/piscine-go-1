package piscine

func IsUpper(s string) bool {
	str := []rune(s)
	for i := 0; i < RuneArrayLength(str); i++ {
		if str[i] > 'Z' || str[i] < 'A' {
			return false
		}
	}
	return true
}
