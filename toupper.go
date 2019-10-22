package piscine

func ToUpper(s string) string {
	str := []rune(s)
	for i := range str {
		if str[i] >= 'a' && str[i] <= 'z' {
			str[i] -= 32
		}
	}
	return string(str)
}
