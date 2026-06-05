package main

import (
	"fmt"
	"log"
	"net"

	"github.com/shradiphylleia/network/internal/request"
)

// deprecating: working with HTTP now
// ----- ----- ---- ----- ----- ----
// func getLinesChannel(f io.ReadCloser) <-chan string {
// 	lines := make(chan string)
// 	go func() {
// 		defer f.Close()
// 		defer close(lines)
// 		acc := ""
// 		for {
// 			data := make([]byte, 8)
// 			count, err := f.Read(data)

// 			if err != nil {
// 				break
// 			}

// 			acc += string(data[:count])
// 			parts := strings.Split(acc, "\n")

// 			for i:=0;i<len(parts)-1;i++ {
// 				lines <- parts[i]
// 			}``
// 			acc = parts[len(parts)-1]
// 		}
// 		if acc != "" {
// 			lines <- acc
// 		}
// 	}()
// 	return lines
// }

func main() {

	// create server: setup listener for tcp
	ln, err := net.Listen("tcp", ":42069")

	if err != nil {
		log.Fatal("hit an err while setting up tcp listener, exiting", err)
	}
	defer ln.Close()
	fmt.Print("Listening on port: 42069")

	for {
		// accept waits for and rtns the next connection
		conn, err := ln.Accept()
		if err != nil {
			log.Fatal("error while setting up tcp listener, exiting", err)
		}
		fmt.Print("connection accepted \n")

		req, err := request.RequestFromReader(conn)

		// deprecating: working with HTTP now
		// ----- ----- ---- ----- ----- ----
		// lines:=  getLinesChannel(conn)
		// single client synchronous server since this runs before next conn is accepted
		// for line := range lines{
		// 	fmt.Println(line)
		// }

		fmt.Print("Request line: \n")
		fmt.Printf("Method: %s\n", req.RequestLine.Method)
		fmt.Printf("Target: %s\n", req.RequestLine.RequestTarget)
		fmt.Printf("Version: %s\n", req.RequestLine.HttpVersion)

		fmt.Println("Headers:")

		for key, value := range req.Headers {
			fmt.Printf("- %s: %s\n", key, value)
		}

	}

}
