package piscine

func ToLower(s string) string {
	str := []rune(s)
	for i := range str {
		if str[i] >= 'A' && str[i] <= 'Z' {
			str[i] += 32
		}
	}
	return string(str)
}
