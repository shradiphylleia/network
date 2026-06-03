package request

import (
	"io"
	"errors"
	"strings"
)

type Request struct {
	RequestLine RequestLine
}

// GET /coffee HTTP/1.1
// /coffee requestTarget 
// method GET

type RequestLine struct {
	HttpVersion   string
	RequestTarget string
	Method        string
}

// Create a request from a reader
// io.Reader: HTTP doesn't care where bytes come from.
func RequestFromReader(reader io.Reader) (*Request, error){

	body, err := io.ReadAll(reader)
	if err!= nil{
		return nil, err
	}

	reqText := string (body)
	lines := strings.Split(reqText, "\r\n")

	requestLine, err :=parseRequestLine(lines[0])

	if err!= nil{
		return nil , err
	}

	return &Request{RequestLine: requestLine} , nil
}

func parseRequestLine(line string) (RequestLine, error){
	components := strings.Split(line, " ")
	if len(components)!= 3 || strings.ToUpper(components[0]) != components[0] || (components[2])!= "HTTP/1.1" {
		return RequestLine{}, errors.New("invalid request line")
	}
		version := strings.Split(components[2], "/")
		return RequestLine{HttpVersion: version[1], RequestTarget: components[1], Method: components[0]} , nil
}
