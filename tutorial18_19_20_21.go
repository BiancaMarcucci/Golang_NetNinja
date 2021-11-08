/* User Input and Switch Statements and Parsing Floats and Saving Files */
// We want the user to be able to choose what to add in the bill //
//intro to Switch statements and on how to convert strings to float64 using strconv.ParseFloat(string,bits)

package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

/* This function will require us to create a new reader everytime we want to ask user for their input,
a better solution is the function "getInput" we define after this function */

func createBill0() bill {
	reader := bufio.NewReader(os.Stdin) // os.Stdin is to get input from terminal specifically
	fmt.Println("Create a new bill name: ")
	name, _ := reader.ReadString('\n') // whatever is captured will be stored as a string, '\n' means "read user input only after they press return key",
	//the placeholder is for error, but we will learn about it later
	// we also want to trim the input, in case before submitting the user puts a lot of white spaces
	name = strings.TrimSpace(name)
	//let's pass name into the function newBill()
	newBill := newBill(name)
	fmt.Println("Created the bill - ", newBill.name)

	return newBill
}

/* getInput function, which takes a string and the pointer to the reader (*bufio.Reader) */
func getInput(question string, reader *bufio.Reader) (string, error) {
	//we want the question to be printed:
	fmt.Println(question)
	input, err := reader.ReadString('\n')

	return strings.TrimSpace(input), err
}

/* let's use getInput in createBill() */
func createBill() bill {
	reader := bufio.NewReader(os.Stdin)
	name, _ := getInput("Create a new bill name: ", reader) // much shorter than in createBill0()
	newBill := newBill(name)
	fmt.Println("Created the bill - ", newBill.name)

	return newBill
}

/*SWITCH STATEMENTS*/
/* Function that asks user to choose an option to do different tasks, Introduction to SWITCH statements */
func promptOptions(b *bill) { // passing the pointer so that i can update tip too!!!!
	reader := bufio.NewReader(os.Stdin)

	option, _ := getInput("Choose option (a - add item, s - save bill, t - add tip):", reader) // for now this only takes an input
	switch option {
	case "a":
		fmt.Println("You chose a")
		name, _ := getInput("What is the item name? ", reader)
		priceString, _ := getInput("What is the item price? ", reader)

		/* to use addItem() function, we need a float64 for the price argument, we thus need to convert it, using strParseFloat(string,bytes)
		but that returns two variables: the produced float and an error, we thus need to select only the float*/

		priceFloat64, err := strconv.ParseFloat(priceString, 64)
		if err != nil {
			fmt.Println("The price must be a number!")
			promptOptions(b)
		}
		fmt.Println(name, priceFloat64)
		b.addItem(name, priceFloat64)
		// and in case user wanted to do more:
		promptOptions(b)
	case "t":
		fmt.Println("You chose t")
		tipString, _ := getInput("What is the tip? ", reader)
		tipFloat64, err := strconv.ParseFloat(tipString, 64)
		if err != nil {
			fmt.Println("The tip must be a number!")
			promptOptions(b)
		}
		fmt.Println(tipFloat64)
		b.updateTip(tipFloat64)
		// and in case user wanted to do more:
		promptOptions(b)
	case "s":
		fmt.Println("You chose save the bill,", b)
		b.saveBill()
		fmt.Println("Bill was saved with name", b.name)
	default: // otherwise:
		fmt.Println("That was not a valid option...")
		promptOptions(b) // if the input is invalid, ask to choose an option again!
	}
}

func main() {
	myBill := createBill()
	promptOptions(&myBill)
	fmt.Println(myBill)
}
