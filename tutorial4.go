/* Printing and Formatting strings  */

package main

import "fmt" // is actually to format strings

func main() {

	/* PRINT */
	// does not print on new line
	fmt.Print("hello,")
	fmt.Print("world!")
	// need escape \n for new line
	fmt.Print("\nhello, ")
	fmt.Print("world!")

	/////Println///////////////////////////

	// does print on new line anything that comes after Println
	fmt.Println("Hi")
	fmt.Println("Bye")

	/////Concatenating////////////////////
	age := 25
	name := "Bianca"
	fmt.Println("Hi, I am", name, "and I am", age, "years old")

	/* FORMATTING */

	// Printf (formatted strings)///////////////////////////////

	// Place holders are eg %v, the general form is %_ (and then add what you need in place of _)
	fmt.Printf("My age is %v and my name is %v \n", age, name)
	fmt.Printf("My age is %q and my name is %q \n", age, name) // works only with strings and puts quotes around variable
	fmt.Printf("age is of type %T \n", age)

	// Sprintf (saves formatted strings)//////////////////////////////
	var myFormattedString = fmt.Sprintf("My age is %v and my name is %v \n", age, name)
	fmt.Print("I saved a formatted string that says: ", myFormattedString)

	// setting precision for floats
	fmt.Printf("My number is %f \n", 22.22)
	fmt.Printf("My number is %0.2f ", 22.22)
}
