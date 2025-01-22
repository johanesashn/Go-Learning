package main

import (
	"fmt"
)

func hello (s string) {
	fmt.Println(s)
}

type person struct {
	name string
	age int
}

// this method is in the person struct
// so it can be accessed by typing person.detail 
// and it would return the person.name
func (p person) detail() string {
	// return p.name
	// can be better like this if you want add another text in the return
	// printF is for printing
	// sprintf is for returning value
	// since we return string with format, then we use sprintf nor printf
	return fmt.Sprintf("The person's name is %s", p.name)
}

type subject struct {
	name string
	teacher person
	// teacher struct {
	// 	name string
	// 	age int
	// }
}

func main(){
	var name string = "johanes"
	var attend bool = false
	var present bool = true
	var age int = 12
	newName := "alberto"
	newAge := 123
	const number = 60 * 60

	var person1 person
	person1.name = "berto"
	person1.age = 12

	var person2 = person {
		name: "johanes",
		age: 12,
	}
	var personName = person2.detail()
	fmt.Println(personName)

	var subject1 = subject {
		name: "computer vision",
		teacher: person {
			name: "albert",
			age: 12,
		},
	}

	
	fmt.Println(person1)
	// with the key of the object
	fmt.Printf("%+v\n", person2)
	fmt.Printf("%+v\n",subject1)

	hello("makan")

	fmt.Println(attend || present)
	fmt.Println("test lagi bang")
	fmt.Println(name)
	fmt.Println(age)
	fmt.Println(newName, " is ", newAge)
	fmt.Println(number)
}