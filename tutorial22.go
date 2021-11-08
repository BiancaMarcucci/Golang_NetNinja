/* Interfaces */
// Interfeces group types together based on their methods,
//if a type has all the methods the interface has, then it can be grouped //

package main

import (
	"fmt"
	"math"
)

/* Shape Interface */
type shape interface {
	// anything that has methods area() and circumf() (both returning float64) then is a "shape"
	area() float64
	circumf() float64
}
type square struct {
	length float64
}
type circle struct {
	radius float64
}

/* square methods */
func (s square) area() float64 {
	return s.length * s.length
}
func (s square) circumf() float64 {
	return s.length * 4
}

/* circle methods */
func (c circle) area() float64 {
	return math.Pi * c.radius * c.radius
}
func (c circle) circumf() float64 {
	return 2 * math.Pi * c.radius
}

/* So now we can pass either circle or square to the printShapeInfo() function!! */
func printShapeInfo(s shape) {
	fmt.Printf("area of %T is: %0.2f \n", s, s.area()) //%T is to print the type of the value
	fmt.Printf("circumference of %T is: %0.2f \n", s, s.circumf())
}

func main() {
	shapes := []shape{
		// because we are defining shapes as a slice containing items of type "shape", we can both have circles and squares!!
		square{length: 15.2},
		circle{radius: 7.5},
		circle{radius: 12.3},
		square{length: 4.9},
	}

	for _, v := range shapes {
		printShapeInfo(v)
		fmt.Println("---")
	}
}
