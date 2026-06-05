package main

import(
	"github.com/shradiphylleia/network/internal/server"
	"log"
	"syscall"
	"os"
	"os/signal"
)

const port = 42069

func main() {
	srv, err := server.Serve(port)
	if err != nil {
		log.Fatalf("Error starting server: %v", err)
	}
	defer srv.Close()
	log.Println("Server started on port", port)

	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)
	<-sigChan
	log.Println("Server stopped")
}