/* Bill Custom Object */
//Our first Struct//
// We create here a bill object that we will then use in tutorial14//

package main

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
