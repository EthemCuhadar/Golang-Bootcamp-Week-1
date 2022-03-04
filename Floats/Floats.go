package main

import (
	"fmt"
	"reflect"
)

func main() {

	var floatVal1 float32 = 15.9
	var floatVal2 float64 = 15.9
	
	// The next line will cause an error because of types.
	//result := floatVal1 == floatVal2
	//fmt.Println(result)
	
	fmt.Println(floatVal1)
	fmt.Println(floatVal2)
	
	fmt.Printf("Type of the floatVal1 is: %v\n", reflect.TypeOf(floatVal1))
	fmt.Printf("Type of the floatVal2 is: %v\n", reflect.TypeOf(floatVal2))
	
}
