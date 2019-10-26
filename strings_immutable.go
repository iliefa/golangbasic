//strings are immutable
//must be converted to a slice of runes

package main

import "fmt"

func mutate(s []rune) string {
	s[0] = 'a' / '' represents a rune
	return string(s)
}

func main() {
	h := "hello"
	fmt.Println(mutate([]rune(h)))
}
