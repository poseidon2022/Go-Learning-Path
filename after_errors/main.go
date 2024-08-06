package main

import (
	"fmt"
)

func index[T comparable] (s []T, element T) int {
	for index, value := range s {
		if value == element {
			return index
		}
	}
	return -1
}

//example let us define a singly linked list

type List[T any] struct {
	next *List[T]
	val T
}
func main() {
	//functions most of the time return an error function along with the operations
	//example is the thing now
	//now to readers

	//type parameters
	//go functions can be made to work on multiple types using the type parameter

	//they have their own syntax
	//appear bn brackets before the functions arguments <=
	//example: a slice of any type T (can be a string or an int)
	mine := []string {
		"This",
		"is",
		"Sparta",
	}
	fmt.Println(index(mine, "kidus")) //this works mind u for types of all

	another := []int {
		10,20,20,
	}
	fmt.Println(index(another, 20))
	//now let us define generic types
	//a type called list. isnt it exciting?

}