package piscine

func CalculateIntegers(array []rune) int {
	count := 0
	for i := range array {
		if array[i] >= '0' && array[i] <= '9' {
			count++
		}
	}
	return count
}

func TrimAtoi(s string) int {
	str := []rune(s)
	answer := 0
	count := CalculateIntegers(str)
	sign := false
	positive := true
	for i := range str {
		if answer == 0 && str[i] == '0' {
			answer = 0
			count--
		} else if answer == 0 && str[i] == '+' {
			if sign == true {
				return 0
			}
			sign = true
		} else if answer == 0 && str[i] == '-' {
			if sign == true {
				return 0
			}
			sign = true
			positive = false
		} else if str[i] >= '0' && str[i] <= '9' {
			ten := 1
			for j := 1; j < count; j++ {
				ten = ten * 10
			}
			count--
			answer = answer + chartoint(str[i])*ten
		}
	}
	if positive == true {
		return answer
	} else {
		return -answer
	}
}
