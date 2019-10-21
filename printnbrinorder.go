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
	} else {
		for i := 1; i <= size(n); i++ {
			ten := 1
			for j := i; j > 1; j-- {
				ten = ten * 10
			}
			z01.PrintRune(rune((n/ten)%10 + 48))
		}
	}
}
