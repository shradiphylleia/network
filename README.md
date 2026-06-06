# Network Systems in Go

Low-level backend systems exploration in Go focused on understanding networking and protocol internals from scratch.


# TCP

Implemented a synchronous TCP echo server using raw sockets and Go's net package.

### Concepts explored
    TCP sockets
    Host/port binding
    Streams vs messages
    Byte buffers
    Partial reads
    Connection lifecycle
    Stream framing
    CLI configuration using flags

### Running the TCP server:

```bash
cd tcp
go run .
```

open another terminal at root:
``` bash
go run client.go 
```

Example: 
go run . --host=0.0.0.0 --port=8080


# HTTP
Currently implementing HTTP/1.1 from scratch to understand:

    request parsing
    headers
    connection handling
    protocol design


Detailed read:

# HTTP From First Principles: RFC 9110 & RFC 9112
A ground-up implementation of HTTP/1.1 in Go built directly on top of raw TCP sockets.

## Architecture

```text
TCP Socket
    │
    ▼
Byte Stream
    │
    ▼
Incremental HTTP Parser
    │
    ├── Request Line
    ├── Headers
    └── Body
    │
    ▼
Request Object
    │
    ▼
Handler
    │
    ▼
Response Writer
    │
    ▼
HTTP Response Serialization
    │
    ▼
TCP Connection
```

---

# Transport Layer Foundations

Before implementing HTTP, the project explored raw TCP and UDP communication.

## TCP

The TCP implementation intentionally reads from connections using extremely small buffers.

```go
buf := make([]byte, 8)
```

This exposes a fundamental networking concept:
> TCP preserves byte ordering, not application-level message boundaries.
A single logical message may arrive fragmented across many reads.
Likewise, multiple logical messages may arrive in a single read.
To handle this correctly, application-level framing must be implemented above TCP.
The server reconstructs logical messages using:
* incremental buffering
* line assembly
* goroutines
* channels

---

## UDP
UDP exploration focused on understanding datagram semantics.
Unlike TCP:
* no connection establishment
* no retransmission
* no ordering guarantees
* packet-oriented communication

Operational testing demonstrated how packets are silently discarded when no listener is available, reinforcing the distinction between stream-oriented and datagram-oriented protocols.

---

# HTTP Request Parsing
HTTP is implemented as a streaming parser rather than a one-shot parser.
The parser accepts arbitrary chunks of bytes and maintains state across reads.
This mirrors how production protocol parsers operate.
---

## State Machine
Request parsing is modeled as a finite-state machine.
```text
Initialized
    │
    ▼
Parsing Request Line
    │
    ▼
Parsing Headers
    │
    ▼
Parsing Body
    │
    ▼
Done
```
Each state consumes only the bytes relevant to that stage.
This allows parsing to continue correctly even when requests arrive fragmented across multiple TCP reads.
---

## Request Line Parsing

Example:
```http
GET /beenon HTTP/1.1
```
Parsed into:

```go
type RequestLine struct {
    Method        string
    RequestTarget string
    HttpVersion   string
}
```

Validation includes:
* method format checks
* HTTP version validation
* request line structure validation
* malformed request detection
---

## Header Parsing
Headers are parsed incrementally, one header at a time.
Example:

```http
Host: localhost:42069
User-Agent: curl/8.0
Accept: */*
```

### Features
#### Case-Insensitive Lookup
The following are treated identically:
```http
Host
HOST
host
```
All keys are normalized internally.

#### Validation
Malformed header names are rejected.
Example:

```http
H©st: localhost
```
results in a parse failure.

#### Multi-Value Header Support
RFC-compliant repeated headers are merged.
Input:

```http
Set-Person: lane
Set-Person: tj
Set-Person: prime
```

Stored as:

```http
set-person: lane, tj, prime
```
---

## Request Body Parsing
Request bodies are framed using:
```http
Content-Length
```
The parser:
* reads body bytes incrementally
* validates expected payload length
* detects oversized payloads
* transitions to completion only after the full body has been received
This introduces explicit message framing at the application layer.
---

# Fragmentation Testing
One common mistake when implementing protocol parsers is assuming an entire request arrives in a single read.
To test correctness, I implemented:

```go
type chunkReader struct
```
which intentionally fragments requests into arbitrary chunk sizes:

```text
1 byte/read
2 bytes/read
3 bytes/read
...
```
This validates parser behavior under realistic TCP fragmentation scenarios.
---
# Response Generation
HTTP responses are serialized manually.
## Status Line
Implemented support for:
```http
HTTP/1.1 200 OK
HTTP/1.1 400 Bad Request
HTTP/1.1 500 Internal Server Error
```
---

## Headers
Default response headers include:

```http
Content-Length
Connection
Content-Type
```
Headers are generated and serialized without framework support.
---

## Response Writer
A custom response abstraction was implemented.
```go
type Writer struct
```
The abstraction provides handler-controlled:

* status codes
* headers
* response bodies
while keeping protocol serialization separate from business logic.
Conceptually similar to:

```go
http.ResponseWriter
```
from Go's standard library.
---

# Concurrent Server Runtime
The server implementation includes:

### Connection Acceptance
```go
net.Listener
```
accept loop running continuously.
### Concurrent Request Handling
Each accepted connection is processed in its own goroutine.

```go
go s.handle(conn)
```

### Graceful Shutdown
Server lifecycle is managed using:
```go
atomic.Bool
```
allowing shutdown without spurious listener errors.
---

# tldr:
* Stream processing over TCP
* Application-level message framing
* Incremental protocol parsing
* Finite-state machine design
* HTTP/1.1 request and response semantics
* Concurrent connection handling
* Transport-agnostic abstractions using `io.Reader` and `io.Writer`
* Protocol correctness under fragmented network reads

HTTP is ultimately just a structured interpretation of a stream of bytes arriving over a TCP connection.

// Everything else is abstraction. //
