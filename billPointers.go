/* Bill Custom Object and its receiver functions with pointers */
// We create here a bill object that we will then use in tutorials 17 to 21//

package main

import (
	"fmt"
	"os"
)

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
		tip:   0.0,
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

/*SAVING DATA TO A FILE ---- using os.WriteFile(nameFile,dataToSave, perm FileMode )
files with perm filemode 0644, will be readable by all the user groups, but writable by the user only.
https://stackoverflow.com/questions/18415904/what-does-mode-t-0644-mean/18415935 */

/* saveBill() */
func (mybill *bill) saveBill() {
	data := []byte(mybill.format()) //we are saving the data as a slice of bytes, after converting to bytees the formatted bill (string)

	err := os.WriteFile(mybill.name+".txt", data, 0644) // if an error occurs, it willl be saved in err
	if err != nil {                                     // if error, panic
		panic(err)
	} // otherwise write file with formatted bill in it
	fmt.Println("Bill was saved to file!")
}
