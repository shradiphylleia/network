package main

import (
	"log"
	"syscall"
	"os"
	"os/signal"
	"github.com/shradiphylleia/network/headers"
	"github.com/shradiphylleia/network/internal/request"
	"github.com/shradiphylleia/network/internal/response"
	"github.com/shradiphylleia/network/internal/server"
)

const port = 42069

func main() {
srv, err := server.Serve(port, func(w *response.Writer, req *request.Request) {

	switch req.RequestLine.RequestTarget {

	case "/kilas":
		w.WriteStatusLine(response.StatusBadRequest)

		w.WriteHeaders(headers.Headers{
			"content-type": "text/html",
		})

		w.WriteBody([]byte(`
			<html>
			<head>
				<title>400 Bad Request</title>
			</head>
			<body>
				<h1>Bad Request</h1>
				<p>tujhko kya fikar</p>
			</body>
			</html>
`))
		return

	case "/buredin":
		w.WriteStatusLine(response.StatusInternalServerError)

		w.WriteHeaders(headers.Headers{
			"content-type": "text/html",
		})

		w.WriteBody([]byte(`
			<html>
			<head>
				<title>500 Internal Server Error</title>
			</head>
			<body>
				<h1>Internal Server Error</h1>
				<p>shakti &amp; kshama</p>
			</body>
			</html>
`))
		return

	default:
		w.WriteStatusLine(response.StatusOK)

		w.WriteHeaders(headers.Headers{
			"content-type": "text/html",
		})

		w.WriteBody([]byte(`
			<html>
			<head>
				<title>[as]</title>
			</head>
			<body>
				<h1>Seedhe maut</h1>
			</body>
			</html>
`))
	}
})

	if err != nil {
		log.Fatalf("Error starting server: %v", err)
	}

	defer srv.Close()

	log.Println("Server started on port", port)

	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)

	<-sigChan

}