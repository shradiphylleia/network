package main

// print library
import (
	"fmt" 
	"math"
 	"strings")


// go primer:

// package level variable
// vars: =: short variable declaration operator, only func outside vars and type , auto infers the type

// constants

const (
	name = "shraddha"
	msg = "no cap"
	day = "friday"
)


func main(){
	var x int = 400
	var y int = 60
	var z int = 4
	// var x,y,z int =400,60,4

	// x, y, z :=400, 60, 4
}

// has pointer as a concept also
// rae strings in backticks
// int float64 bool string var ptr *int [nil]
// Printf, Sprintf- formatted output

// type aliases


// module is the unit of dependency management and project org similar to package.json , pom.xml(Maven)
// module: project
// packgae: directory of go files
// installing dependencies : go get <> subsequently updates the module file [ go.md , go.sum]

// ! go.sum is depenedency verification file, storing cryptographic checksums of the exact dependency versions of the project
// hash of needed packages
// go.sum is not lockfile



// os package : interact with os platform independent interface 
// open files , mkdir read env vars , get process info , delete files, handle stdin/stdout
// type-assert errors *os.PathError , extract structured info from the error

// if we create something we MUST use it.
// unsused variables as compile errors
// go automatically inserts semicolons at line endings so else must be on the same line as the closing of the if block
// idiomatic go usually avoid else: err path exists early , success path clean and less nesting
// The methods of File correspond to file system operations. All are safe for concurrent use. The maximum number of concurrent operations on a File may be limited by the OS or the system. The number should be high, but exceeding it may degrade performance or cause other issues.


// file gave hash: location ? : different value
// how to get params?
// *file: different values


// reading bytes: 
// []byte : slice of bytes
// make would create an underlying array and slice pointing to it
// data := make ([]byte, 8)
// slice is dynamic view of bytes [] sequence of bytes
// go zero initializes memory by default, each element is a byte (uint8) so 0 = 00000000 in binary

// only for 
// for {} infinte loop