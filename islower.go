package piscine

func IsLower(s string) bool {
	str := []rune(s)
	for i := 0; i < RuneArrayLength(str); i++ {
		if str[i] > 'z' || str[i] < 'a' {
			return false
		}
	}
	return true
}
