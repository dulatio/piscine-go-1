package piscine

import "github.com/01-edu/z01"
import "fmt"

func PrintComb2() {
	for i := 0; i <= 98; i++ {
		for j := i + 1; j <= 99; j++ {
			fmt.Readln(i / 10)
			fmt.Readln(i % 10)
			z01.PrintRune(' ')
			fmt.Readln(j / 10)
			fmt.Readln(j % 10)
			z01.PrintRune(',')
			z01.PrintRune(' ')
		}
	}
	z01.PrintRune('\n')
}
