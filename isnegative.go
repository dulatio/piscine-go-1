package piscine

import "github.com/01-edu/z01"

func IsNegative(nb int) rune {

	if nb >= 0 {
		return 'F'
	} else {
		return 'T'
	}
}