package main

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/mock"
)

// func testCalculate(t *testing.T) {
// 	if Calculate(2) != 4 {
// 		t.Error("expected to be four")
// 	}
// }

//using testify, it looks something like this

//import assert
//t here is just a context and all reports are made to this object when runnning
//the test

//this is used for record keeping and easy logging of records

type MyStruct struct {}

type MyStructInterface interface {
	DoAmazingStuff(int)		(bool, error)
}

func (m *MyStruct) DoAmazingStuff(input int) (bool, error) {
	fmt.Print("This is amazing") //this is some check or smth quite often
	return true, nil
}


func AmazingFunction(m MyStructInterface) {
	m.DoAmazingStuff(123)
}

type MyStructMocked struct {
	mock.Mock
}

func (m *MyStructMocked) DoAmazingStuff(input int) (bool, error) {
	args := m.Called(input)
	return args.Bool(0), args.Error(1)
}

func testAmazingFunction(t *testing.T) {
	testObj := new(MyStructMocked)

	testObj.On("DoAmazingStuff", 123).Return(true, nil)
	AmazingFunction(testObj)
	testObj.AssertExpectations(t)
}