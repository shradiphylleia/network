package main

import ("fmt" 
		"os"
 		"log"
		"strings")

func main(){

	// open
	file, err := os.Open("messages.txt")
	if err != nil {
	log.Fatal (err)
	}
	fmt.Println("file opened", file.Name())

	// read : while err not show up, infinite for loop
	acc := ""
	for{
	// buffer :fresh	
	data := make ([]byte, 8)
	// file object stores current position
	count , err := file.Read(data)
	if err != nil{
		break
	}
	
	acc = acc +string(data[:count])
	parts := strings.Split(acc, "\n")
	
	if len(parts)>1{
		fmt.Printf("%s \n", parts[0])
		acc = parts[len(parts)-1]
	}
	}
}

