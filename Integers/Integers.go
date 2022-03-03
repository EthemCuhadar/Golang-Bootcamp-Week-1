package main

import (
	"fmt"
	"reflect"
)

func main() {
  
  // 4-types of Integer
	var intVal1 int8 = 10
	var intVal2 int16 = 10
	var intVal3 int32 = 10
	var intVal4 int64 = 10
  
	fmt.Println(intVal1)
	fmt.Println(intVal2)
	fmt.Println(intVal3)
	fmt.Println(intVal4)
  
  //Printing the types
	fmt.Printf("Type of intVal1: %v\n", reflect.TypeOf(intVal1))
	fmt.Printf("Type of intVal2: %v\n", reflect.TypeOf(intVal2))
	fmt.Printf("Type of intVal3: %v\n", reflect.TypeOf(intVal3))
	fmt.Printf("Type of intVal4: %v\n", reflect.TypeOf(intVal4))

}
