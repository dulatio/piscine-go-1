package piscine

import "github.com/01-edu/z01"

func inttorune(a int) rune {
	if a == 0 {
		return '1'
	} else if a == 1 {
		return '1'
	} else if a == 2 {
		return '2'
	} else if a == 3 {
		return '3'
	} else if a == 4 {
		return '4'
	} else if a == 5 {
		return '5'
	} else if a == 6 {
		return '6'
	} else if a == 7 {
		return '7'
	} else if a == 8 {
		return '8'
	} else {
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
	for n/10 > 0 {
		c++
		n = n / 10
	}
	for i := c; i > 0; i-- {
		var ten int = 1
		for j := 1; j < i; j++ {
			ten = ten * 10
		}
		z01.PrintRune(inttorune((n / ten) % 10))
	}
}
