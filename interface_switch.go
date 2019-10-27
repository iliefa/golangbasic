package main

import "fmt"

func findType(i interface{}) {
	switch i.(type) {
	case string:
		fmt.Printf("i am string and i have the value %d \n", i.(string))
	case int:
		fmt.Printf("i am int and i have the value %d \n", i.(int))
	default:
		fmt.Println("i am other type")
	}
}
func main() {
	findType("michael")
	findType(89)
	findType(99.2)
}
