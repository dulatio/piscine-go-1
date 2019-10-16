package piscine

func chartoint(c rune) int {
	if c == '0' {
		return 0
	} else if c == '1' {
		return 1
	} else if c == '2' {
		return 2
	} else if c == '3' {
		return 3
	} else if c == '4' {
		return 4
	} else if c == '5' {
		return 5
	} else if c == '6' {
		return 6
	} else if c == '7' {
		return 7
	} else if c == '8' {
		return 8
	} else {
		return 9
	}
}

func BasicAtoi(s string) int {
	str := []rune(s)
	answer := 0
	for i := range str {
		if answer == 0 && str[i] == '0' {
			answer = 0
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
