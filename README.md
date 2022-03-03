# Golang

![Golang Image]()

-------------------------------------------------------------------

## General Information

- **Golang**, aka "**Go**", is an open source, statically typed and compiled programming language designed at Google by ***Robert Griesemer***, ***Rob Pike*** and ***Ken Thompson***. 

- Go was publicly announced in Novermber 2009. 

- General usage areas of the language are ***Cloud Computing*** and ***Web Applications and Services***.

--------------------------------------------------------

## Getting started with "Hello World"

Our first program will print "Hello World". The code is shown below. 

```go
package main

import "fmt"

fun main(){
    ftm.Println("Hello World")

}
```

To run the code, type "go run filename.go" in terminal. 

```console
go run hello-world.go
Hello World
```

--------------------------------------------------------------------

# Starting with Variable Types

In Go, "var" declares variables.

--------------------------------------------------------

## Booleans

Boolean data type is the data type which can be ***True*** or ***False***. Example code is shown below.

```go
package main

import "fmt"


fun main(){

    var boolVal1 bool
    fmt.Println(boolVar1)

    var boolVal2 = true
    fmt.Println(boolVal2)
}
```

```console
go run Booleans.go
false
true
```

* Default bool value in Go is false
* When boolVal2 is pointed to ***true***, it is not necessary to decleare type bool.

-------------------------------------------------------

## Integers

There are 4 types of integers in Go which are **int8**, **int16**, **int32** and **int64**.

```go
package main

import (
    "fmt"
    "reflect"
)

func main() {

    var intVal1 int8 = 10
    var intVal2 int16 = 10

    fmt.Println(intVal1)
    fmt.Println(intVal2)

    fmt.Printf("Type of intVal1: %v\n", reflect.TypeOf(intVal1))
    fmt.Printf("Type of intVal2: %v\n", reflect.TypeOf(intVal2))

}
```

```console
10
10
Type of intVal1: int8
Type of intVal1: int16
```

In Go programming language, variable types are very strict. Therefore, you cannot apply basic operations for different types of variables such as adding and multiplication. You can see an example of this situation below with decleared values of ***intVal1*** and ***intVal2***.

```go
fmt.Println(intVal1 + intVal2)
```

```console
# command-line-arguments
.\Integers.go:25:22: invalid operation: intVal1 + intVal2 (mismatched types int8 and int16)
```

* The general types of these 2 values may be the same. However, there are also 4 types in the integers. Hence, integer types are different for this situation.

--------------------------------------------


