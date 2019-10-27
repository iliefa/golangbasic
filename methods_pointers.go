package main

import "fmt"

type rectangle struct {
	length int
	width  int
}

func perimeter(r *rectangle) {
	fmt.Println("perimeter is", 2*(r.length+r.width))
}

func (r *rectangle) perimeter() {
	fmt.Println("perimeter is", 2*(r.length+r.width))
}

func main() {
	r := rectangle{10, 5}
	p := &r //pointer to r
	perimeter(p)
	p.perimeter()
	// perimter(r) will not work
	r.perimeter() // interpreted as (&r).perimeter()
}
