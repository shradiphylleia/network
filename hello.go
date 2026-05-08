package main

// print library
import ("fmt" "math" "strings")


// go primer:

// package level variable


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



