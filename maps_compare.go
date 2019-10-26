package main

import (
	"fmt"
	"reflect"
)

//compare 2 maps
func main() {
	//can not be compared using == ,== can be used to check if map is nil

	map1 := map[string]int{
		"one":   1,
		"two":   2,
		"three": 3,
	}

	map2 := map[string]int{
		"one":   1,
		"two":   2,
		"three": 4,
	}

	fmt.Println(reflect.DeepEqual(map1, map2))
}
