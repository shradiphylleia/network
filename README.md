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