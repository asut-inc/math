package main

import "fmt"
import "math"

type Circle struct {
	x float64
	y float64
	r float64
}

func (c *Circle) area() float64 {
    return math.Pi * c.r*c.r
}

type Rectangle struct {
	x float64
	y float64
	w float64
	h float64
}

func (r *Rectangle) area() float64 {
    return r.w*r.h
}

func main(){
	var c Circle
	c = Circle{x:0, y:0, r:5}
	r := Rectangle{x:0, y:0, w:5, h:6}

	fmt.Println(c.area())
	fmt.Println(r.area())
}