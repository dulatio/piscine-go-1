package piscine

func Len(str string) int {
	length := 0
	strc := []rune(str)
	for index := range strc {
		length = index
	}
	return length
}

func StrRev(s string) string {
	str := []rune(s)
	answer := []rune(s)
	j := 0
	for i := Len(s); i >= 0; i-- {
		answer[j] = str[i]
		j++
	}
	return string(answer)
}
