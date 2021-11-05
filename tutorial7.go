/* Loops  */
// For in Go is used for what in other languages would be while, for and for in //

package main

import "fmt"

func main() {
	/* WHILE equivalent */
	// want to loop through some block of code while a variable x is below a certain value----> ~while loop
	x := 0
	for x < 5 {
		fmt.Println("The value of x is: ", x)
		x++
	}
	/* FOR equivalent */
	for i := 0; i < 5; i++ {
		fmt.Println("The value of i is: ", i)
	}
	//looping through a slice:
	names := []string{"mario", "luigi", "yoshi", "peach"}
	for i := 0; i < len(names); i++ {
		fmt.Println("The current name is: ", names[i])
	}

	/* FOR IN equivalent */
	// we need the 'range' keyword, used to initialise and update both the index and value of the slice
	for index, value := range names {
		fmt.Printf("The current position is %v and the value %v \n", index, value)
	}
	// if we do not want to use the index, you can just use a place holder ('_')
	for _, value2 := range names {
		fmt.Printf("The current name is  %v \n", value2)
	}

}
