package piscine

func IsAlpha(s string) bool {
	str := []rune(s)
	for i := 0; i < RuneArrayLength(str); i++ {
		if (str[i] > 'z' || str[i] < 'a') && (str[i] > '9' || str[i] < '0') && (str[i] > 'Z' || str[i] < 'A') {
			return false
		}
	}
	return true
}
