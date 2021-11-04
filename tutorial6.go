// Few extra packages examples, to be used with strings and slices etc//
package main

import (
	"fmt"
	"sort"
	"strings" // this contains many methods to be used with strings
	// to sort elements in slices and user-defined collections
)

func main() {
	//Strings///////////////////////////////////////////////
	greetings := "hello there friends!"
	fmt.Println(strings.Contains(greetings, "Ginger"))        // This will return false
	fmt.Println(strings.ReplaceAll(greetings, "hello", "hi")) // this will replace all "hello" with "hi"---NB it returns a NEW STRING
	fmt.Println(strings.ToUpper(greetings))                   // this uppercases the passed string ---NB it returns a NEW STRING
	fmt.Println(strings.Index(greetings, "e"))                // to find at what index you find first occurrance of "e"

	// the original greetings is not changed:
	fmt.Println("Original string is still: ", greetings)

	// Split is a method that returns a slice containing parts of the string, devided accordingly to what declared
	fmt.Println(strings.Split(greetings, " ")) // Will split the string at every space

	//Sort/////////////////////////////////////////////////
	ages := []int{20, 25, 30, 44, 21, 12, 56, 70}

	sort.Ints(ages) // This sorts a slice of integers, will permanently change the original slice
	fmt.Println(ages)

	indexInSortedSlice := sort.SearchInts(ages, 30) // will return the index of first occurrance of an element in the SORTED slice
	fmt.Println(indexInSortedSlice)

	// Sort alphabetically slice of strings (NB. the capitals)----Changes permanently the original slice
	name := []string{"bianca", "Anna", "Maria"} // Capitals have priority on not capital letters, do Maria will be come before bianca
	sort.Strings(name)
	fmt.Println(name)

	// SearchStrings, returns index
	fmt.Println(sort.SearchStrings(name, "Anna"))

}
