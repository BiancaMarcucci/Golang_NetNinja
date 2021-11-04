// Arrays & Slices //

package main

import "fmt"

func main() {

	//ARRAYS///////////////////////////////

	//Ways to declare Arrays. NB the length is declared and fixed, the type of the variables in the array is also defined.
	var ages [3]int = [3]int{20, 25, 30}
	var ages2 = [3]int{10, 20, 30}
	ages3 := [4]int{1, 2, 3, 4}
	//Can also use method len() to get the length of the array
	fmt.Println(ages, ages2, ages3, len(ages3))

	//SLICES//////////////////////////////

	// Similar to arrays, but lenght is NOT FIXED
	var scores = []int{100, 20, 76}
	scores[2] = 70
	// Append returns a new array and thus needs to be saved in a variable (or update the one you just appended to)
	var scoresUpdated = append(scores, 200)
	fmt.Println(scores)
	fmt.Println(scoresUpdated)

	//Slices ranges
	rangeOne := scoresUpdated[1:3] // this creates a new array that takes a slice of scoresUpdated, from position 1 (included) to position 3 (excluded)
	rangeTwo := scoresUpdated[1:]  // if end limit is not specified, it will take everything until end of array/slice (last item included)
	fmt.Println(rangeOne, rangeTwo)
}
