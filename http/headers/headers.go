package headers

import (
	"errors"
	"strings"
	"unicode"
)

type Headers map[string]string
func NewHeaders() Headers {
	return make(Headers)
}
func (h Headers) Parse(data []byte) (n int, done bool, err error) {
	idx := strings.Index(string(data), "\r\n")
	//need more data
	if idx == -1 {
		return 0, false, nil
	}
	//end of header
	if idx == 0 {
		return 2, true, nil
	}
	line := string(data[:idx])
	parts := strings.SplitN(line, ":", 2)
	if len(parts) != 2 {
		return 0, false, errors.New("invalid header")
	}
	//leading/trailing whitespace in key
	if parts[0] != strings.TrimSpace(parts[0]) {
		return 0, false, errors.New("invalid header")
	}
	key := parts[0]
	value := strings.TrimSpace(parts[1])

	for _, r := range key {
		if !(unicode.IsLetter(r) ||
			unicode.IsDigit(r) ||
			r == '-') {
			return 0, false, errors.New("invalid header")
		}
	}
	key = strings.ToLower(key)
	existingValue, exists := h[key]
	if exists {
		h[key] = existingValue + ", " + value
	} else {
		h[key] = value
	}
	return idx + 2, false, nil
}

func (h Headers) Get(key string) string {
	return h[strings.ToLower(key)]
}