// [3:01 PM] Mullapudi, Srikanth (Unverified)
// Find the most repeated item in an Array
//
//	Input: {3, "a", "a", "a", 2, 3, "a", 3, "a", 2, 4, 9, 3}
//
//	Output: a -> 5 times
package main

import (
	"fmt"
	"os"
)

func main() {

	var input []interface{} = []interface{}{4, "a", "a", "a", 2, 3, "a", 3, "a", 2, 4, 9, 3}

	item, count, err := findRepeated(input)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	fmt.Printf("Item (%+v) is repeated %d times", item, count)
}

// findRepeated would find the item that is most repeated in an
// array of interface
// return the first most repeated, not all
func findRepeated(arr []interface{}) (interface{}, int, error) {

	itemap := map[interface{}]int{}
	var maxitem interface{}
	var curmax int = 0

	for _, item := range arr {
		c := itemap[item] // check to see if the item is there in the map already
		c += 1
		if c > curmax {
			curmax = c
			maxitem = item
		}
		itemap[item] = c
	}
	fmt.Printf("Map: %+v\n", itemap)

	return maxitem, curmax, nil
}

