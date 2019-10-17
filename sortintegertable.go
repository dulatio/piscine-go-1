package piscine

func TabLen(table []int) int {
	length := 0
	for range table {
		length++
	}
	return length
}

func SortIntegerTable(table []int) {
	change := false
	for i := 0; i < TabLen(table)-1; i++ {
		if table[i] > table[i+1] {
			temp := table[i]
			table[i] = table[i+1]
			table[i+1] = temp
			change = true
		}
	}
	if change == true {
		SortIntegerTable(table)
	}
}
