package piscine

func BasicAtoi(s string) int {
	str := []rune(s)
	answer := 0
	for i := 0; i <= Len(s); i++ {
		if answer == 0 && str[i] == '0' {
			answer = 0
		} else {
			ten := 1
			for j := i; j < Len(s); j++ {
				ten = ten * 10
			}
			answer = answer + (int(str[i])-48)*ten
		}
	}
	return answer
}
