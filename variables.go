package main
import (
		"fmt"
		"math"
)
func main(){
	var age int=29 // var age =29
	var width, height= 50,100
	fmt.Println("my age is",age)
	fmt.Println("width is ",width,"height is",height)
	//if not declared will be 0
	var weight int
	fmt.Println("weight is",weight)
	weight=100
	fmt.Println("weight is ",weight)

	//to declare multiple variables
	var(
		name="andrei"
		newage = 20
		heights int 
	)
	fmt.Println("name is",name,"age ",newage,"height",heights)


	// short declaration: name := initialvalue 
	//when at least one of the variables in the left side of := is newly declared. 
	sex, music := "female", "madonna"
	fmt.Println("i am",sex,"i listen",music)

	a,b := 10,20.0
	fmt.Println(a,b)
	b,c :=30.2,40.1
	fmt.Println(b,c)
	d := math.Min(b,c)
	fmt.Println(d)
}