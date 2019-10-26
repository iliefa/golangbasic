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
	area,_ := rectProps(0.10, 15)
	//we want to know only the area
	fmt.Printf("area is %f",area)
}