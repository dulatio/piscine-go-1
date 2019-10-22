package piscine

func Capitalize(s string) string {
	str := []rune(s)
	if str[0] >= 'a' && str[0] <= 'z' {
		str[0] -= 32
	}
	for i := 1; i < RuneArrayLength(str); i++ {
		if str[i] >= 'a' && str[i] <= 'z' && ((str[i-1] > '9' || str[i-1] < '0') && (str[i-1] > 'z' || str[i-1] < 'a') && (str[i-1] > 'Z' || str[i-1] < 'A')) {
			str[i] -= 32
		} else if str[i] >= 'A' && str[i] <= 'Z' && ((str[i-1] <= '9' && str[i-1] >= '0') || (str[i-1] >= 'a' && str[i-1] <= 'z') || (str[i-1] <= 'Z' && str[i-1] >= 'A')) {
			str[i] += 32
		}
	}
	return string(str)
}
