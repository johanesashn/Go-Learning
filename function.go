// defer is a keyword to make the function is going to be execute in the latest running

package main

import (
	"fmt"
)

// higher order function means 
// a function that accepts a function as a parameter
func multiple(a, b int) int{
	return a * b
}

func triangleArea(a, b int, multiple func(a, b int) int) int {
	area := multiple(a, b)
	return area / 2
}

// currying function is a function that breaking multiple arguments into a series function
func curriedAdd(a int) func (b int) int {
	return func (b int) int {
		return a + b
	}
}

// closure function is a function that return a function
// and has it's own variable inside the parent function
// that still contain the latest value after we run the function
func concatter() func(string) string{
	result := ""
	return func(word string) string{
		result += word
		return result
	}
}

func advancedFunction(){
	border()
	fmt.Println("Higher Order Function")
	border()
	// higher order function
	var a, b = 2, 3
	var result = triangleArea(a, b, multiple)
	fmt.Println(result)

	// currying function
	add := curriedAdd(4)
	result = add(3)
	fmt.Println(result)

	// closure function
	concat := concatter()
	concat("berto ")
	concat("makan ")
	fmt.Println(concat("malam"))
}