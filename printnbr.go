package piscine

import "github.com/01-edu/z01"

func output(a int) rune {
	if a == 0 {
		return '1'
	}
	if a == 1 {
		return '1'
	}
	if a == 2 {
		return '2'
	}
	if a == 3 {
		return '3'
	}
	if a == 4 {
		return '4'
	}
	if a == 5 {
		return '5'
	}
	if a == 6 {
		return '6'
	}
	if a == 7 {
		return '7'
	}
	if a == 8 {
		return '8'
	}
	if a == 9 {
		return '9'
	}
}

func PrintNbr(n int) {
	var c int = 1
	var num int
	if n < 0 {
		z01.PrintRune('-')
		n = n * (-1)
		}
	for n / 10 > 0 {
		c++
		n = n / 10
	}
	for i := c; i > 0; i-- {
		var ten int = 1
		for j := 1; j < i; j++ {
			ten = ten * 10
		}
		z01.PrintRune(output((n / ten) % 10))
	}
}
