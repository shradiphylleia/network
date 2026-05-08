package main

import ("fmt" 
		"os"
 		"log")

func main(){

	// open
	file, err := os.Open("messages.txt")
	if err != nil {
	log.Fatal (err)
	}
	fmt.Println("file opened", file.Name())

	// read : while err not show up
	for{
	data := make ([]byte, 8)
	count , err := file.Read(data)
	if err != nil{
		break
	}
	fmt.Printf( "read: %s\n", string(data[:count]))

	}
}

