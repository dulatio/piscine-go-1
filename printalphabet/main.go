package main

import "fmt"

func main() {

	var letter = []rune{'a', 'b', 'c', 'd', 'e', 'f', 'g', 'h', 'i', 'j', 'k', 'l', 'm', 'n', 'o', 'p', 'q', 'r', 's', 't', 'u', 'v', 'w', 'x', 'y', 'z'}

	// convert rune slice to string
	var alphabet = string(letter)

	fmt.Printf("%v\n", alphabet)

	
}