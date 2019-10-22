package piscine

func Join(strs []string, sep string) string {
	result := ""
	for _, word := range strs {
		result = result + word + sep
	}
	array := []rune(result)
	array[RuneArrayLength(array)-1] = 127
	return string(array)
}
