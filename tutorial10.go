/* Multiple returns values */
// Want a function that take a string (full name) and returns Both Initials (if present) but Capital

package main

import (
	"fmt"
	"strings"
)

//NB how we specify both returns:
func getInitials(fullName string) (string, string) {
	capitalFullName := strings.ToUpper(fullName) //Capitalise whole string
	names := strings.Split(capitalFullName, " ") // get an array containing first and last name (if present), separating the string at space

	var initials []string // Initialising an empty slice to which we will append the initials
	for _, v := range names {
		initials = append(initials, v[:1])
	}
	// if both first and last name ---> return both initials
	if len(initials) > 1 {
		return initials[0], initials[1]
	}
	// if only one initial
	return initials[0], "_"
}

func main() {
	firstName1, lastName1 := getInitials("Snow White")
	firstName2, lastName2 := getInitials("Lisa")
	fmt.Println(firstName1, lastName1)
	fmt.Println(firstName2, lastName2)
}
