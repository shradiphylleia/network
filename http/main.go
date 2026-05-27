package main

import (
	"fmt"
	"io"
	"log"
	"os"
	"strings"
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
	file, err := os.Open("messages.txt")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("file opened:", file.Name())
	lines := getLinesChannel(file)
	for line := range lines {
		fmt.Printf("read: %s\n", line)
	}
}