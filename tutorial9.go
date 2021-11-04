//Using Functions//
package main

import (
	"fmt"
	"math"
)

var names []string = []string{"mario", "luigi", "yoshi", "peach", "bowser"}

// Functions that do not return anything but print something://////////////////////
func sayGreeting(myString string) {
	fmt.Printf("Good morning %v! \n", myString)
}
func sayBye(myString string) {
	fmt.Printf("Bye %v! \n", myString)
}

//Functions that take more than a parameter, one of which is a function itself://////////////////////
func cycleNames(mySlice []string, myFunc func(string)) {
	for _, v := range names {
		myFunc(v)
	}
}

//Functions that return something (NB you need to specify the type of what will be returned)
func circleArea(radius float64) float64 {
	return math.Pi * radius * radius
}

func main() {
	// we will use the functions we just created outside the main function:
	cycleNames(names, sayGreeting)
	cycleNames(names, sayBye)
	a1 := circleArea(20.55)
	a2 := circleArea(1)
	fmt.Printf("Area 1 is: %0.3f , and Area 2 is: %0.3f \n", a1, a2)
}
