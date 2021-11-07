/* Structs & Custom Types */
// While other OOP languages have classes, Go uses Structs to create Custom types//
// We will create a custom object for a "bill" ----> see bill.go file//

package main

import "fmt"

func main() {
	myBill := newBill("Mario's bill")
	fmt.Println(myBill)

	/* Adding tip and menu items */

	myBill.items = map[string]float64{"bread": 2.5, "pasta alla matriciana": 12.8, "water": 2.3}
	myBill.tip = 3
	fmt.Println(myBill)
}
