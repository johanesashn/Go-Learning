// NOTE 
// variadic (...var) it is like args in function js, can get many params
// spread (var...) like ...var in js, just to spread the value

package main

import "fmt"

func array_learning() {
	// they are imported from main.go
	// because they have same package
	// that's why they don't need to be imported
	border()
	fmt.Println("this is array")
	border()

	// array has a fixed size
	// var numbers [3]int
	numbers2 := [3]int{1,2,3}
	fmt.Println(numbers2)

	// instead of using array, commonly we use slice
	// because can be shrinked or grown
	var slice []int
	slice = append(slice, 2, 3)
	fmt.Println(slice)

	var matrix = make([][]int, 0)
	// same with the bellow
	// var matrix [][]int
	// matrix := [][]int{}
	for i := 0; i < 3; i++{
		var row []int
		for j := 0; j < 3; j++{
			row = append(row, j)
		}
		matrix = append(matrix ,row)
	}

	fmt.Println(matrix)

	// too loop the slice
	// for index, element := range slice
	for index, row := range matrix {
		fmt.Println(index, row)
	}
}