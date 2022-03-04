package main

import (
	"fmt"
	"reflect"
)

func main() {

	var strVal1 string = "patika"
	var strVal2 string = "bootcamp"
	
	
	fmt.Println(strVal1)
	fmt.Println(strVal2)
	
	fmt.Printf("Type of the strVal1 is: %v\n", reflect.TypeOf(strVal1))
	fmt.Printf("Type of the strVal2 is: %v\n", reflect.TypeOf(strVal2))
	
	fmt.Println(strVal1 == strVal2, " if they are equal")
	fmt.Println(reflect.TypeOf(strVal1) == reflect.TypeOf(strVal2), " if they are same type")
	
	fmt.Println(strVal1 + " " + strVal2)
}
