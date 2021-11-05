/* Pass-by-value  */
// Non-pointer values and Pointer wrapper values //

package main

import "fmt"

//no return function
func doesNotUpdateString(stringWillNotChange string) {
	stringWillNotChange = "new string"
}

//return function
func updateString(stringWillChange string) string {
	stringWillChange = "updated string"
	return stringWillChange
}

//to update maps
func updateMap(myMap map[string]int) {
	myMap["d"] = 4
}

/////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

func main() {

	/*NON POINTER VALUES

	Non-pointer values are: strings,integers,floats,booleans,arrays & structs... */

	//When any of these is passed into a function, the function creates a new copy of the variable in a new memory block,
	// that is different from the passed-in variable location. Thus, any change made on the variable INSIDE the function,
	// is a change made to the COPY and not the original variable.

	name := "string still needs changing"
	doesNotUpdateString(name) // this will change stringWillNotChange, NOT name
	fmt.Println(name)         //This will print "string still needs changing" and not "new string"

	//To actually change string name//////////////////////////////
	name = updateString(name) // save what updateString(name) returns in string name
	fmt.Println(name)         //This will print "updated string"

	/*POINTER WRAPPER VALUES

	Pointer values are: slices,maps,functions.......*/

	//When any of these is passed into a function, the function creates a new copy of the POINTER in a new memory block,
	// that is different from the passed-in variable location. BUT that copy will point to the SAME VARIABLE VALUE,
	// so any change made inside the function, will be applied to the original variable too!
	alphaNum := map[string]int{
		"a": 1,
		"b": 2,
		"c": 3,
		"d": 44,
	}
	fmt.Println(alphaNum) // The value of "d" is still 44.
	updateMap(alphaNum)
	fmt.Println(alphaNum) // Now the value of "d" will be 4.
}
