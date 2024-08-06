package main

import (
	"fmt"
	"math"
)

//interface definition

type another interface{}
type Abser interface {
	Abs() float64 //float64 is the return type, can only define methods inside an interface 
}

type I interface {
	M() //no return 
}

type T struct {
	Name string
}

type Ver struct {
	X,Y float64
}

func (v *Ver) Abs() float64 { //a pointer to of type vert implements the abs function
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func (t T) M() {
	fmt.Println(t.Name)
}

func (v *Ver) Scale(f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

type Person struct {
	name string
	age int
}

func (p Person) String() string { //this acts as a default printing method for the type Person
	return fmt.Sprintf("(%v, %v) is the age and the name of the person", p.name, p.age)
}

func Scale(v *Ver, f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}
func main() {
	//manipulating objects using pointers
	mine := Ver{3,4}
	mine.Scale(10)
	fmt.Println(mine)


	//pointers and functions, instead of the above method. notice the difference 
	Scale(&mine, 10)
	fmt.Println(mine)
	//method and pointer indirection
	//notice the difference here
	//try passing a non-pointer to a function taking a pointer as an argument => compile error
	//but methods taking a pointer as a receiver can also take values other than pointers
	//here is an example
	mine.Scale(10) //this took a receiver of type ver and not *ver
	//works just fine
	//Scale(mine, 10) => this doesnt work tho since the argument must be a receiver for a function

	//we use pointer reveivers for one of the two reasons
	//to avoid copying values everytime a method is called
	//to manipulate the POINTED AT value directly

	//an interface type
	//a set of method signatures
	//a value of interface type can hold any value that implements the methods specified by the interface 

	//let us see some examples
	var a Abser //an abser interface 
	//now let us see who we can assign to a
	v := Ver{3,4}
	//we can not do that since a ver type bu itself doesnot implement the abs function
	//but we can assign a pointer to it, since the pointer impllements the abs function
	a = &v
	fmt.Println(a.Abs()) //this is the same as calling the abs method with a vert type
	fmt.Println(v.Abs())

	//example for an interface implemented implicitly
	var e I = T{"Kidus Melaku"}
	e.M()

	//interface values
	//under the hood the interfaces hold a tuple of values and type
	//let us print the value and the type for the above assigned interface
	fmt.Printf("(%v, %T)\n",e,e) //value with its type

	//we can also have an empty interface for this case
	var i interface{} //this can hold any value in the future and the type assignation is also automatic
	fmt.Printf("(%v, %T)\n",i,i) //both the value and the type are undeclared or ni.
	i = 42
	fmt.Printf("(%v, %T)\n",i,i)

	//type assertions
	//we can not directly call the method on an interface without the underlying value
	//there is also an intetace that specifies zero methods. any value can take that
	//interface since every value implements at least 0 methods

	//example
	var d another
	d = 33
	fmt.Printf("(%v, %T)\n", d,d) //the same as interface i okay?

	var f interface{} = "hello" //this is a string as u can see. if we try to give it a type other than a string, this will trigger a panic
	//example
	t, ok := f.(string)
	fmt.Println(t, ok)
	m,ok := f.(float64) //if there is an ok to it, it will be false and a default 0 value will be given as a value
	fmt.Println(m, ok)
	//but without the ok, it will cause a panic
	//the whole process is called type assertion

	//type switches, combo of type and switches
	switch _,nk := f.(float32); nk{ //just like if assignemnt and statements
		case false:
			fmt.Println("this is not possible")
		case true:
			fmt.Println("this doesnt make any sense what so ever")
	}

	var kidus interface{}
	kidus = "hello"
	do(kidus)
	do(true)
	//lemme show you a weird combo now, I dunno why we use it hulla tbh
	//see the function do below
	//now let us delve a bit up into stringers

	fmt.Println("Amisu") //<p>amisu</p>
	//one of the first things that an fmt function looks for in an application is the presence of a stringer function
	//it will take the format of the stringer function whenever trying to print out anything

	amesyas := Person{
		"Amesyas Taddese",
		12,	
	}
	fmt.Println(amesyas)

	//the string function is quite important tf
	//see the exercise on stringers
}

func do(i interface{}) {
	switch v := i.(type) {
		case int:
			fmt.Printf("%v is an intege\n", v)
		case string:
			fmt.Println(len(v),"nothing to be ashamed of darling")
		case bool:
			fmt.Println("this is of type lie")
	}
}