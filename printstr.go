package piscine

import "github.com/01-edu/z01"

func PrintStr(str string) {
	strc := []rune(str)
	for _, letter := range strc {
		z01.PrintRune(letter)
	}
}
