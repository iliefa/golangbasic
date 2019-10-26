package main

import (
	"fmt"
)

func rectProps(length,width float64) (area,perimeter float64){
	area = length * width
	perimeter = (length + width) * 2
	return 
}

func main(){
	area,perimeter := rectProps(0.10, 15)
	fmt.Printf("area is %f , perimter is %f",area,perimeter)
}