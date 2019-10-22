package piscine

func Index(s string, s1 string) int {
	str := []rune(s)
	toFind := []rune(s1)
	var res int
	var found bool
	for i := 0; i <= RuneArrayLength(str)-RuneArrayLength(toFind); i++ {
		found = true
		res = i
		for j := 0; j < RuneArrayLength(toFind); j++ {
			if str[i+j] != toFind[j] {
				found = false
				break
			}
		}
		if found {
			break
		}
	}
	if found {
		return res
	} else {
		return -1
	}
}
