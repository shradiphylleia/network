package request

import (
	"errors"
	"io"
	"strings"
)

const (
	stateInitialized = iota
	stateDone
)

type Request struct {
	RequestLine RequestLine
	state       int
}

type RequestLine struct {
	HttpVersion   string
	RequestTarget string
	Method        string
}

func RequestFromReader(reader io.Reader) (*Request, error) {
	body, err := io.ReadAll(reader)
	if err != nil {
		return nil, err
	}
	req := &Request{
		state: stateInitialized,
	}
	_, err = req.parse(body)
	if err != nil {
		return nil, err
	}
	return req, nil
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
	// consumed request-line+CRLF
	return requestLine, idx+2, nil
}

func (r *Request) parse(data []byte) (int, error) {
	switch r.state {
	case stateInitialized:
		requestLine, consumed, err := parseRequestLine(data)
		if err != nil {
			return 0, err
		}
		// parser needs more bytes
		if consumed == 0 {
			return 0, nil
		}
		r.RequestLine = requestLine
		r.state = stateDone
		return consumed, nil
	case stateDone:
		return 0, nil
	}
	return 0, nil
}