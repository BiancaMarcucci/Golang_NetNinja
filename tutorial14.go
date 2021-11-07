/* Pointers */
// How to use pointers as arguments to access the original underlying data within functions//
// We can manually create pointers for non-pointer values //
// & is used to get the a pointer to the memory location of the variable in front of which we put & //
// to de-reference a pointer (getting the value of the variable at the memory address it addresses to) we use * //

package main

import "fmt"

func doNotUpdateName(x string) {
	x = "Mario"
}

func doUpdateName(x *string) { // this accepts a POINTER TO A STRING
	*x = "Mario" // this CHANGES THE STRING AT MEMORY ADDRESS
}

func main() {
	name := "Luigi"

	/* FUNCTION DOES NOT UPDATE THE STRING */

	doNotUpdateName(name)
	fmt.Println("Memory address of name is: ", &name) // &name will give the address of variable name="Luigi"

	//Pointers can be stored in variables!
	myPointer := &name
	fmt.Println("Memory address : ", myPointer) // myPointer is the same as &name

	// To de-reference myPointer --> *myPointer
	fmt.Println("Value stored at memory address : ", *myPointer)
	fmt.Println(name)

	/* FUNCTION DO UPDATE THE STRING */

	doUpdateName(myPointer) //NEED POINTER
	fmt.Println("Memory address of name is: ", myPointer)
	fmt.Println("Value stored at memory address is now: ", *myPointer) //Mario
	fmt.Println(name)
	fmt.Println(name == *myPointer)

}
