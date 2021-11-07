/* Bill Custom Object and its receiver functions with pointers */
// We create here a bill object that we will then use in tutorial17//

package main

import "fmt"

type bill struct {
	name  string
	items map[string]float64
	tip   float64
}

/* Function to generate a new bill object */

func newBill(name string) bill {
	b := bill{
		name:  name,
		items: map[string]float64{}, // empty map by default
		tip:   0,
	}
	return b
}

/* RECEIVER FUNCTIONS */

/* format() */
// a function that is associated only to this custom type bill
// function to format the bill into a string

func (mybill *bill) format() string { /* NB how for using pointers you only need to ass * and not need to dereference it later (GO does it for you) */
	formattedBill := "Bill breakdown: \n"
	var total float64 = 0 // could have just used the tip to be fair, but we can add it later, just to follow the videos.

	//looping through menu items:
	for dish, price := range mybill.items {
		formattedBill += fmt.Sprintf("%-25v ...$%0.2f \n", dish+":", price) // %-25v means "add characters to the left so that the variable is long 25 characters"
		total += price
	}
	//total (also adding the tip)
	formattedBill += fmt.Sprintf("%-25v ...$%0.2f \n", "tip:", mybill.tip)
	formattedBill += fmt.Sprintf("%-25v ...$%0.2f \n", "total:", total+mybill.tip)

	return formattedBill
}

/* updateTip() */
// a function that update the tip to the passed value

func (mybill *bill) updateTip(newTip float64) {
	// (*mybill).tip=newTip... but Go does the de-referencing for you, so you can write :
	mybill.tip = newTip
}

/* addItem() */
// a function to add a new item and its price to the bill

func (mybill *bill) addItem(dish string, price float64) {
	mybill.items[dish] = price
}
