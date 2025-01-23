package main 

import (
	"fmt"
)

type shape interface {
	// the struct would be considered as shape if they use all the method on this interface
	// in this interface there is only getArea method, that's why every struct that has getArea method
	// is considered as shape.
	getArea() int
}

type rectangle struct {
	width, height int
}

type square struct {
	length int
}

func (r rectangle) getArea() int {
	return r.width * r.height
}

func (s square) getArea() int {
	return s.length * s.length
}

// _ there is no param used inside the function that's why just use _
func getDetail(_ shape) {
	fmt.Println("this is shape")
}

func learning() {
	fmt.Println("start of the program")
	var r = rectangle {
		width: 12, 
		height: 12,
	}
	var s = square {
		length: 10,
	}
	fmt.Println(r.getArea())
	fmt.Println(s.getArea())
	getDetail(r)
	getDetail(s)
}