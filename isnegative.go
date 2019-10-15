package main

import piscine ".."

func IsNegative(nb int) rune {

	if nb >= 0 {
		return 'F'
	} else {
		return 'T'
	}
}

func main() {

	piscine.IsNegative(1)
	piscine.IsNegative(0)
	piscine.IsNegative(-1)

}
