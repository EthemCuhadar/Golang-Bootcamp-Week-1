# Golang

-------------------------------------------------------------------

## General Information

- **GOlang**, aka "**Go**", is an open source, statically typed and compiled programming language designed at Google by ***Robert Griesemer***, ***Rob Pike*** and ***Ken Thompson***. 

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


