package main

import (
	"fmt"
	"io"
	"log"
	"strings"
	"net"
)

func getLinesChannel(f io.ReadCloser) <-chan string {
	lines := make(chan string)
	go func() {
		defer f.Close()
		defer close(lines)
		acc := ""
		for {
			data := make([]byte, 8)
			count, err := f.Read(data)

			if err != nil {
				break
			}

			acc += string(data[:count])
			parts := strings.Split(acc, "\n")

			for i:=0;i<len(parts)-1;i++ {
				lines <- parts[i]
			}
			acc = parts[len(parts)-1]
		}
		if acc != "" {
			lines <- acc
		}
	}()
	return lines
}


func main() {
	
	// create server: setup listener for tcp
	ln , err := net.Listen("tcp", ":42069")
	
	if err!= nil{
		log.Fatal("hit an err while setting up tcp listener, exiting",err)
	}
	defer ln.Close()
	fmt.Print("Listening on port: 42069")

	for{
		// accept waits for and rtns the next connection
		conn,err:= ln.Accept()
		if err!=nil{
			log.Fatal("error while setting up tcp listener, exiting",err)
		}
		fmt.Print("connection accepted \n")
		lines:=  getLinesChannel(conn)

		// single client synchronous server since this runs before next conn is accepted
		for line := range lines{
			fmt.Println(line)
		}
	}
	
}