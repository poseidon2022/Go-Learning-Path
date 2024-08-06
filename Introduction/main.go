package main

// in go unlike python, the submodule is called with a slash example math/rand not math.rand
import (
	"fmt"
	"math"
)

type Vertex struct {
	X,Y float64
}

func (v Vertex) abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func Abs(v Vertex) float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func main() {
	s := []struct {
		isStupid bool
		age int
	} {
		{isStupid: true, age: 13},
	}
	fmt.Println(s)

	//creating slices with make, example, make can take 3 arguments, the slice type, length and cap
	//example

	another := make([]int, 0, 6) //dynamic property unlike arrays and normal slices
	for i := 0;i < 6; i ++ {
		another = append(another, i)
	}
	fmt.Println(len(another), cap(another))
	another = append(another, 8)
	fmt.Println(len(another), cap(another))

	board := [][]int { //we can also add the types for the inner slices if we want to but it doesnt matter
		{1,2,3},
		{4,5,6},
		{7,8,9},
	}
	fmt.Println(board)

	//and the append function can add more than one element at a time
	board = append(board, []int{1,2,3,}, []int{3,4,5})
	fmt.Print(board,"\n")

	//now on maps
	mine := map[string]int {
		"kidus": 11,
		"jona": 22,
	}

	fmt.Println(mine) //ok returns false if the element is in the map and false otherwise, but it will give it a default value of 0
	value, ok := mine["kidus"]
	fmt.Println(value, ok)
	anoda, okk := mine["yanet"]
	fmt.Println(anoda, okk)

	x,y := adder(), adder()
	for i:=0; i<10; i++ {
		fmt.Println(x(i), y(2*i))
	}
	//passing functions as arguments
	// u can not directly define a function inside another function
	//but this is possible
	// anota := func(x int, y int) int {
	// 	return x*y
	// }

	//methods on go
	//a function with a special receiver argument
	myPoint := Vertex{3,4}
	fmt.Println(myPoint.abs()) //the abs function can also be weitten as a regular function. not a method
	fmt.Println(Abs(myPoint))
}

func adder() func(int) int {
	sum := 0
	return func(x int) int {
		sum += x
		return sum
	}
}
func compute(fn func(float64, float64) float64) float64 {
	return fn(3, 4)
}