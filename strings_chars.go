// rune is alias of int32 and represents a unicode point in Go

package main

import (
	"fmt"
	"unicode/utf8"
)

func printCharsandBytes(s string) {
	for index, rune := range s {
		fmt.Printf("%c starts and byte %d \n", rune, index)

	}
}

func length(s string) {
	fmt.Printf("length of %s is %d \n", s, utf8.RuneCountInString(s))
}
func main() {
	name := "Señorita"
	printCharsandBytes(name)

	// length of string:
	length(name)
}
