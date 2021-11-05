/* HOW TO DECLARE VARIABLES IN GO */

package main

import "fmt"

func main() {

	/* STRINGS */
	// ways to declare Strings
	var nameOne string = "mario"
	var nameTwo = "luigi"
	// to declare an empty string:
	var nameThree string

	fmt.Println(nameOne, nameTwo, nameThree)

	// changing third name to a different string
	nameThree = "peach"
	fmt.Println(nameOne, nameTwo, nameThree)

	// another way to declare variables without the "var", is using :=
	nameFour := "yoshi"
	fmt.Println(nameFour)

	/* INTEGERS */
	var ageOne int = 20
	var ageTwo = 30
	ageThree := 40

	fmt.Println(ageOne, ageTwo, ageThree)

	/* BITS & MEMORY */
	// you can choose how many bits the numbers will take, eg an integer that takes 8 bits is int8 (from -128 to 127)

	var num8Bits int8 = 120
	fmt.Println(num8Bits)

	// to have NOT NEGATIVE integers "un-signed"
	var numUint uint = 20
	// can mix with bits notation:
	var numUint82 uint = 21 // can go from 0 to 255
	fmt.Println(numUint, numUint82)

	/* FLOATS */
	// either float32 or float64, the latter being the defgault if := is used

	var scoreOne float32 = 33.33
	var scoreTwo float64 = 66.66
	scoreThree := 111.11 // this is a float64

	fmt.Println(scoreOne, scoreTwo, scoreThree)
}
