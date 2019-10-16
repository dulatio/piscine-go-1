package piscine

func Atoi(s string) int {
	str := []rune(s)
	answer := 0
	sign := false
	positive := true
	for i := range str {
		if answer == 0 && str[i] == '0' {
			answer = 0
		} else if answer == 0 && str[i] == '+' {
			if sign == true {
				return 0
			}
			sign := true
		} else if answer == 0 && str[i] == '-' {
			if sign == true {
				return 0
			}
			sign := true
			positive := false
		} else if str[i] < '0' || str[i] > '9' {
			return 0
		} else {
			ten := 1
			for j := i; j < Len(s); j++ {
				ten = ten * 10
			}
			answer = answer + chartoint(str[i])*ten
		}
	}
	if positive == true {
		return answer
	} else {
		return -answer
	}
}
