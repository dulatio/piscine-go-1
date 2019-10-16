package piscine

func StrLen(str string) int {
	length := 0
	strc := []rune(str)
	for index := range strc {
		length = index + 1
	}
	return length
}
