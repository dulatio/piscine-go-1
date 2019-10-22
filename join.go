package piscine

func Join(strs []string, sep string) string {
	result := ""
	var index int
	for i := range strs {
		index = i
	}
	for i, word := range strs {
		result = result + word
		if i != index {
			result = result + sep
		}
	}
	return result
}
