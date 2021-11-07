/* Structs & Custom Types */
// We will create a custom object for a "bill" ----> see billPointers.go file//

package main

import "fmt"

func main() {
	myBill := newBill("Peach's bill")
	fmt.Println(myBill)

	/* Adding tip and menu items */

	myBill.addItem("bread", 2)
	myBill.addItem("pasta alla matriciana", 13)
	myBill.addItem("water", 2)
	myBill.updateTip(5)
	fmt.Println(myBill.format())
}
