package main

import (
	"flag"
	"log"
	"github.com/shradiphylleia/network/tcp/config"
	"github.com/shradiphylleia/network/tcp/server"
)


var cfg config.Config
// server < cmd line configs

func setupFlags(){
	// take cli arg "host" and store it in config.Host
	// --host= 
	// default to 0.0.0.0 if value not provided and description
	flag.StringVar(&cfg.Host, "host", "0.0.0.0", "host for the server")
	flag.IntVar(&cfg.Port, "port", 8080 , "port for the server")
	flag.Parse()
}

func main(){
	setupFlags()
	log.Println("been on")
	server.RunSyncTCP(cfg)
}