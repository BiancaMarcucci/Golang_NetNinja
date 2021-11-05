/*Package Scope*/
// How to use variables and functions from different files from the same package//
//could have used any file here, but let's use greetings.go, we will use variables and functions from it,
// viceversa we will also define a variable (at root level, outside the main() function) in tutorial11.go and we will use it in greetings.go

package main

import "fmt"

// we can declare a variable outside the main function to make it visible to greetings.go
var score = 99.5

func main() {
	//We use sayHello() and we loop through slice 'points' from greetings.go
	sayHello("Vision")

	for _, point := range points {
		fmt.Println(point)
	}
	// using another function from greetings.go
	showScore()

}
