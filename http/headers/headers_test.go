package headers


import (
	"testing"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	)


func TestHeadersParse(t *testing.T) {
// Test: Valid single header
headers := NewHeaders()
data := []byte("Host: localhost:42069\r\n\r\n")
n, done, err := headers.Parse(data)
require.NoError(t, err)
require.NotNil(t, headers)
assert.Equal(t, "localhost:42069", headers["host"])
assert.Equal(t, 23, n)
assert.False(t, done)

// Test: Invalid spacing header
headers = NewHeaders()
data = []byte("       Host: localhost:42069\r\n\r\n")
n, done, err = headers.Parse(data)
require.Error(t, err)
assert.Equal(t, 0, n)
assert.False(t, done)

// Test
headers = NewHeaders()
data = []byte("HOST: localhost:42069\r\n\r\n")
n, done, err = headers.Parse(data)

require.NoError(t, err)
assert.Equal(t, "localhost:42069", headers["host"])
assert.Equal(t, 23, n)
assert.False(t, done)

// Test
headers = NewHeaders()
data = []byte("H©st: localhost:42069\r\n\r\n")

n, done, err = headers.Parse(data)
require.Error(t, err)
assert.Equal(t, 0, n)
assert.False(t, done)

// Test
headers = NewHeaders()
headers["set-person"] = "shraddha"

data = []byte("Set-Person: shraddha-sharma\r\n\r\n")

n, done, err = headers.Parse(data)

require.NoError(t, err)
assert.Equal(t, "shraddha, shraddha-sharma", headers["set-person"])
assert.False(t, done)
assert.Equal(t, 29, n)

}