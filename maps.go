package main

import (
	"fmt"
)

func learning_maps() {
	border()
	fmt.Println("This is maps")
	border()

	// this is the first way
	// ages := make(map[string]int)
	// ages["berto"] = 21

	// second way to make map
	ages := map[string]int{
		"berto": 21,
	}

	fmt.Println(ages)

	// we use rune commonly when it store 1 char, 
	// instead of using string
	// run will change the char into the number
	test := make(map[rune]int)
	// it works, becasue the initial of the map is going to be 0
	test['a'] += 1
	fmt.Println(test)

	testAgain := make(map[rune]string)
	fmt.Println("value is", testAgain['a'])
}