package request

import (
	"errors"
	"io"
	"strings"
	"github.com/shradiphylleia/network/headers"
)

const (
	requestStateInitialized = iota
	requestStateParsingHeaders
	requestStateDone
)

type Request struct {
	RequestLine RequestLine
	Headers     headers.Headers
	state       int
}

type RequestLine struct {
	HttpVersion   string
	RequestTarget string
	Method        string
}

// deprecated -- parse
// func RequestFromReader(reader io.Reader) (*Request, error) {
// 	body, err := io.ReadAll(reader)
// 	if err != nil {
// 		return nil, err
// 	}
// 	req := &Request{
// 		state: stateInitialized,
// 	}
// 	_, err = req.parse(body)
// 	if err != nil {
// 		return nil, err
// 	}
// 	return req, nil
// }

func RequestFromReader(reader io.Reader) (*Request, error) {
	req := &Request{
		state:  requestStateInitialized,
		Headers: headers.NewHeaders(),
	}
	buf := make([]byte, 1024)
	for {
		n, err :=reader.Read(buf)
		if err !=nil {
			return nil, err
		}
		remaining := buf[:n]
		for len(remaining) >0 {
			consumed, err :=req.parse(remaining)
			if err != nil {
				return nil, err
			}
			if consumed==0 {
				break
			}
			remaining=remaining[consumed:]

			if req.state== requestStateDone {
				return req, nil
			}
		}
	}
}
func parseRequestLine(data []byte) (RequestLine, int, error) {
	idx := strings.Index(string(data), "\r\n")
	// need more data
	if idx == -1 {
		return RequestLine{}, 0, nil
	}
	line := string(data[:idx])
	components := strings.Split(line, " ")
	if len(components) != 3 ||
		strings.ToUpper(components[0]) != components[0] ||
		components[2] != "HTTP/1.1" {
		return RequestLine{}, 0, errors.New("invalid request line")
	}
	version := strings.Split(components[2], "/")
	requestLine := RequestLine{
		HttpVersion:   version[1],
		RequestTarget: components[1],
		Method:        components[0],
	}
	return requestLine, idx+2, nil
}

func (r *Request) parseSingle(data []byte) (int, error) {
	switch r.state {
	case requestStateInitialized:
		requestLine, consumed, err := parseRequestLine(data)
		if err != nil {
			return 0, err
		}
		if consumed == 0 {
			return 0, nil
		}
		r.RequestLine = requestLine
		r.state = requestStateParsingHeaders
		return consumed, nil

	case requestStateParsingHeaders:
		consumed, done, err := r.Headers.Parse(data)
		if err != nil {
			return 0, err
		}
		if done {
			r.state = requestStateDone
		}
		return consumed, nil
	case requestStateDone:
		return 0, nil
	}

	return 0, nil
}

func (r *Request) parse(data []byte) (int, error) {

	totalBytesParsed := 0

	for r.state != requestStateDone {

		n, err := r.parseSingle(data[totalBytesParsed:])

		if err != nil {
			return 0, err
		}

		if n == 0 {
			break
		}

		totalBytesParsed += n
	}

	return totalBytesParsed, nil
}