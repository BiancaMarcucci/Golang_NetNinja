/* Bill Custom Object */
//Our first Struct//
// We create here a bill object that we will then use in tutorial15_16//

package main

import "fmt"

type bill1 struct {
	name  string
	items map[string]float64
	tip   float64
}

/* Function to generate a new bill object */

func newBill1(name string) bill1 {
	b := bill1{
		name:  name,
		items: map[string]float64{}, // empty map by default
		tip:   0,
	}
	return b
}

/* Receiver Function */
// a function that is associated only to this custom type bill
// function to format the bill into a string

func (mybill bill1) format() string {
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
