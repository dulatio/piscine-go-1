package piscine

import "github.com/01-edu/z01"

func size(n int) int {
	res := 0
	for ; n > 0; n /= 10 {
		res++
	}
	return res
}

func PrintNbrInOrder(n int) {

	if n == 0 {
		z01.PrintRune('0')
	}

	var array [19]int

	for i := 0; i <= 18; i++ {
		array[i] = 20
	}

	for i := 0; i < size(n); i++ {
		ten := 1
		for j := i; j > 0; j-- {
			ten = ten * 10
		}
		array[i] = (n / ten) % 10
	}

	SortIntegerTable(array[:])

	for i := 0; i < size(n); i++ {
		z01.PrintRune(rune(array[i] + 48))
	}
}
