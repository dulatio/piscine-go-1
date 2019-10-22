package piscine

func IsPrintable(s string) bool {
	str := []rune(s)
	for i := 0; i < RuneArrayLength(str); i++ {
		if str[i] > 127 || str[i] < 32 {
			return false
		}
	}
	return true
}
