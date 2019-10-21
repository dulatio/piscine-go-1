package piscine

import "github.com/01-edu/z01"

var array [8]int

func threatened(queen_nb int, row int) bool {
	for i := 0; i < queen_nb; i++ {
		if array[i] == row || array[i] == row-queen_nb+i || array[i] == row+queen_nb-i {
			return true
		}
	}
	return false
}

func solution(current int) {
	if current == 8 {
		for i := 0; i <= 7; i++ {
			z01.PrintRune(rune(array[i] + 49))
		}
		z01.PrintRune(10)
	} else {
		for i := 0; i <= 7; i++ {
			if threatened(current, i) == false {
				array[current] = i
				solution(current + 1)
			}
		}
	}
}

func EightQueens() {
	solution(0)
}
