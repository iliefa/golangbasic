package main
import (
	"fmt"
)

func main(){
	// type conversion does not work
	i :=55
	j := 2.1
	sum := float64(i) +j
	fmt.Println(sum)
}