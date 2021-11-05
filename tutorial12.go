// Maps //
// Similar to Dictionaries in Python or Objects on Javascript //
// But the keys and values can be of any type, as long as all the keys are the same and the values are all the same //
// Syntax is----> mapName := map[keysType]valuesType{...} //
package main

import "fmt"

func main() {

	// Creating a menu, in which the keys are strings and the values are floats:
	menu := map[string]float64{
		"soup":           4.99,
		"pie":            7.99,
		"salad":          6.99,
		"toffee pudding": 3.55, // the comma is needed also for the last pair
	}

	// Map that has Integer keys and String values:
	luckyNum := map[int]string{
		2: "mario",
		6: "peach",
		8: "luigi",
	}

	// We can either print the map itself, or any of its value (passing the corresponding key)
	fmt.Println(menu)
	fmt.Println(luckyNum)

	fmt.Println(menu["pie"])
	fmt.Println(luckyNum[2])

	//We can change the values of a map:
	luckyNum[2] = "bowser"

	fmt.Println(luckyNum)
	fmt.Println(luckyNum[2])

	// Looping maps
	for key, value := range menu {
		fmt.Println(key, "-", value)
	}
}
