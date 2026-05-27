// udp listener
package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
)

func main() {

	udpAddress, err := net.ResolveUDPAddr("udp", "localhost:42069")
	if err != nil {
		log.Fatal("err while establishing udp listener", err)
	}

	conn, err := net.DialUDP("udp", nil, udpAddress)

	if err != nil {
		log.Fatal("err while Dial: connects to a server", err)
	}
	defer conn.Close()
	// convert keybaord input to stream
	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Print(">")
		// read until newline encountered
		text, err := reader.ReadString('\n')
		if err != nil {
			log.Fatal(err)
		}

		_, err = conn.Write([]byte(text))
		if err != nil {
			log.Fatal(err)
		}

	}

}
