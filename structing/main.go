package main

import (
	"fmt"
)

type myStruct struct {
	intField int
}

func (ms myStruct) addByValue(x int) {
	///it is OK to leave the return type off a function or method if we are not returning a value
	ms.intField += x
	fmt.Println("ByValue internal value ", ms.intField)
}

func (ms *myStruct) addByReference(x int) {
	ms.intField += x
	fmt.Println("ByReference internal value ", ms.intField)
}

func main() {
	myVar := myStruct{1}
	myPtr := &myStruct{2}

	myVar.addByValue(3)
	fmt.Println("main func myVar value ", myVar)
	myVar.addByReference(3)
	fmt.Println("main func myVar value ", myVar)
	fmt.Println("\n")

	myPtr.addByValue(3)
	fmt.Println("main func myPtr value ", myPtr)
	myPtr.addByReference(3)
	fmt.Println("main func myPtr value ", myPtr)

}
