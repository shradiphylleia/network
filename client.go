package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
)

func main() {
	conn, err := net.Dial("tcp", "localhost:8080")
	if err != nil {
		panic(err)
	}
	defer conn.Close()
	fmt.Println("connected to server")
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Print("> ")
		text, _ := reader.ReadString('\n')
		conn.Write([]byte(text))
		msg, _ := bufio.NewReader(conn).ReadString('\n')

		fmt.Println("echo:", msg)
	}
}