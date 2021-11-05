/* Booleans & Conditionals  */

package main

import "fmt"

func main() {
	names := []string{"mario", "luigi", "yoshi", "peach", "bowser"}

	for index, value := range names {
		// continue is to keep looping
		if index == 1 {
			fmt.Println("Continuing at position ", index)
			continue
		}
		// break to break completely from looping
		if index > 2 {
			fmt.Println("Breaking the loop at position ", index)
			break
		}
		fmt.Printf("The value at position %v is %v \n", index, value)
	}
}
