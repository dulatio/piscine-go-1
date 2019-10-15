package piscine

import "github.com/01-edu/z01"

func PrintComb2() {
	for i := 0; i <= 98; i++ {
		for j := i + 1; j <= 99; j++ {
			var a = i / 10;
			var b = i % 10;
			var c = j / 10;
			var d = j % 10;
			z01.PrintRune('a')
			z01.PrintRune('b')
			z01.PrintRune(' ')
			z01.PrintRune('c')
			z01.PrintRune('d')
			z01.PrintRune(',')
			z01.PrintRune(' ')
		}
	}
	z01.PrintRune('\n')
}
