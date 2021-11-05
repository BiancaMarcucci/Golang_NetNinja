/*For Tutorial 11 on Package Scope (could have done with any other file, as we all put them in package main)
we will define variables and functions we will then use in tutorial11.go and will also use a variable from tutorial11.go */

package main

import "fmt"

var points = []int{20, 90, 100, 45, 70}

func sayHello(myString string) {
	fmt.Println("hello", myString)
}

// we use variable we declared in tutorial11.go and then we will use showScore() in tutorial11.go
func showScore() {
	fmt.Println("You scored this many points: ", score)
}
