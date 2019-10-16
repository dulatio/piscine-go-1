package piscine

func BasicAtoi2(s string) int {
	str := []rune(s)
	answer := 0
	for i := range str {
		if answer == 0 && str[i] == '0' {
			answer = 0
		} else if str[i] < '0' && str[i] > '9' {
			return 0
		} else {
			ten := 1
			for j := i; j < Len(s); j++ {
				ten = ten * 10
			}
			answer = answer + chartoint(str[i])*ten
		}
	}
	return answer
}
