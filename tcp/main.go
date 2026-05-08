package main

import (
	"flag"
	"log"
	"fmt"
)

func setupFlags(){
	flag.StringVar(&config.Host, "host", "0.0.0.0", "host for the server")
	flag.IntVar(&config.Port, "port", 4646 , "port for the server")
}

func main(){
	log.Println("been on")

}