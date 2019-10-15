package main

import "fmt"
import "github.com/01-edu/z01"


func main() {

	var alphabet string = "abcdefghijklmnopqrstuvwxyz"
	
	for i := 0; i <= 25; i++ {
		z01.PrintRune(rune(alphabet[i]))
	}

}
