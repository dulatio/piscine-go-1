package piscine

func BasicJoin(strs []string) string {
	result := ""
	for _, word := range strs {
		result = result + word
	}
	return result
}
