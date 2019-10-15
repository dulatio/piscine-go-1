package isnegative

import "github.com/01-edu/z01"

func IsNegative(nb int) rune {

	if nb >= 0 {
		return 'F'
	} else {
		return 'T'
	}
}

func main() {

	z01.PrintRune(IsNegative(1))
	z01.PrintRune('\n')
	z01.PrintRune(IsNegative(0))
	z01.PrintRune('\n')
	z01.PrintRune(IsNegative(-1))
	z01.PrintRune('\n')

}
